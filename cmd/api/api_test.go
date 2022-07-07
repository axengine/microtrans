package api

import (
	"context"
	"errors"
	"github.com/axengine/go-saga"
	"github.com/axengine/go-saga/storage/kafka"
	"github.com/axengine/go-saga/storage/redis"
	"github.com/axengine/utils/id/uuid"
	"go-micro.dev/v4/client"
	"log"
	"microtrans/config"
	"microtrans/proto/order"
	"microtrans/proto/wallet"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	zkAddrs        = []string{"192.168.10.32:2181"}
	brokerAddrs    = []string{"192.168.10.32:9092"}
	partitions     = 1
	replicas       = 1
	returnDuration = 50 * time.Millisecond
)

func TestCreateOrder(t *testing.T) {
	cfg := config.DefaultConfig
	oClient := order.NewOrderService(cfg.Service.OrderSrvName, client.DefaultClient)
	resp, err := oClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
		OrderId: uuid.GenUUID(),
		Goods:   1,
		Price:   21.23,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestSetOrder(t *testing.T) {
	cfg := config.DefaultConfig
	oClient := order.NewOrderService(cfg.Service.OrderSrvName, client.DefaultClient)
	resp, err := oClient.SetOrder(context.Background(), &order.SetOrderStatusRequest{
		OrderId: "12",
		Status:  0,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestPayment(t *testing.T) {
	cfg := config.DefaultConfig
	wClient := wallet.NewWalletService(cfg.Service.WalletSrvName, client.DefaultClient)
	resp, err := wClient.Payment(context.Background(), &wallet.PaymentRequest{
		Uid:     1,
		OrderId: "12",
		Value:   12.34,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestCompensatePayment(t *testing.T) {
	cfg := config.DefaultConfig
	wClient := wallet.NewWalletService(cfg.Service.WalletSrvName, client.DefaultClient)
	resp, err := wClient.CompensatePayment(context.Background(), &wallet.CompensatePaymentRequest{
		Uid:     1,
		OrderId: "12",
		Value:   12.34,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestMicroTrans(t *testing.T) {
	cfg := config.DefaultConfig

	uid := int64(1)
	orderId := uuid.GenUUID()
	goodsId := int32(2)
	price := 12.34

	// 创建订单
	oClient := order.NewOrderService(cfg.Service.OrderSrvName, client.DefaultClient)
	if _, err := oClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
		OrderId: orderId,
		Uid:     uid,
		Goods:   goodsId,
		Price:   price,
	}); err != nil {
		t.Fatal(err)
	}

	wClient := wallet.NewWalletService(cfg.Service.WalletSrvName, client.DefaultClient)

	// action函数必须幂等
	payFunc := func(ctx context.Context, uid int64, orderId string, value float64) error {
		log.Println("exec payFunc")
		//return errors.New("mock error")
		_, err := wClient.Payment(ctx, &wallet.PaymentRequest{
			Uid:     uid,
			OrderId: orderId,
			Value:   value,
		})
		if err != nil {
			log.Println("Payment err:", err)
		}
		return err
	}

	// 补偿函数必须幂等,如果返回错误，会一直重试;所以对服务要求很高，不能陷入死循环
	compensatePayFunc := func(ctx context.Context, uid int64, orderId string, value float64) error {
		log.Println("exec compensatePayFunc")
		_, err := wClient.CompensatePayment(ctx, &wallet.CompensatePaymentRequest{
			Uid:     uid,
			OrderId: orderId,
			Value:   value,
		})
		if err != nil {
			log.Println("CompensatePayment err:", err)
		}
		return err
	}

	setOrderFunc := func(ctx context.Context, orderId string, status int32) error {
		log.Println("exec setOrderFunc")
		if status == 1 {
			return errors.New("mock err")
		}
		//panic("mock panic")
		_, err := oClient.SetOrder(ctx, &order.SetOrderStatusRequest{
			OrderId: orderId,
			Status:  status,
		})
		if err != nil {
			log.Println("setOrderFunc err:", err)
		}
		return err
	}

	compensateSetOrderFunc := func(ctx context.Context, orderId string, status int32) error {
		log.Println("exec compensateSetOrderFunc")
		_, err := oClient.SetOrder(ctx, &order.SetOrderStatusRequest{
			OrderId: orderId,
			Status:  status - 1,
		})
		if err != nil {
			log.Println("compensateSetOrderFunc err:", err)
		}
		return err
	}

	//store, err := kafka.NewKafkaStorage(zkAddrs, brokerAddrs, partitions, replicas, returnDuration, saga.LogPrefix,
	//	log.New(os.Stdout, "saga_", log.LstdFlags))
	//if err != nil {
	//	t.Fatal(err)
	//}
	logPrefix := "saga_"
	store, err := redis.NewRedisStore("192.168.10.16:6379", "111111", 14, 2, 5, logPrefix)
	if err != nil {
		t.Fatal(err)
	}

	sec := saga.NewSEC(store, logPrefix)
	sec.AddSubTxDef("钱包支付", payFunc, compensatePayFunc)
	sec.AddSubTxDef("修改订单状态", setOrderFunc, compensateSetOrderFunc)

	// sagaID 要求全局不冲突，在事务未结束（正常、异常）时不能重复使用sagaID
	var sagaID string = "1"

	// ctx贯穿所有action的执行
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = sec.StartSaga(ctx, sagaID).
		ExecSub("钱包支付", uid, orderId, price).
		ExecSub("修改订单状态", orderId, int32(1)).EndSaga()
	// err是action执行的err
	if err != nil {
		log.Println("最终执行结果:", err)
	}

	// 如果saga内部日志操作失败会直接panic，所以对kafka要求很高
	// 如果使用内存存储日志，程序崩溃后无法恢复
}

func BenchmarkMicroTrans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = microtrans(strconv.FormatInt(int64(i), 10))
	}
}

func microtrans(sagaID string) error {
	cfg := config.DefaultConfig

	uid := int64(1)
	orderId := uuid.GenUUID()
	goodsId := int32(2)
	price := 12.34

	// 创建订单
	oClient := order.NewOrderService(cfg.Service.OrderSrvName, client.DefaultClient)
	if _, err := oClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
		OrderId: orderId,
		Uid:     uid,
		Goods:   goodsId,
		Price:   price,
	}); err != nil {
		return err
	}

	wClient := wallet.NewWalletService(cfg.Service.WalletSrvName, client.DefaultClient)

	// action函数必须幂等
	payFunc := func(ctx context.Context, uid int64, orderId string, value float64) error {
		//return errors.New("mock error")
		_, err := wClient.Payment(ctx, &wallet.PaymentRequest{
			Uid:     uid,
			OrderId: orderId,
			Value:   value,
		})
		if err != nil {
			log.Println("Payment err:", err)
		}
		return err
	}

	// 补偿函数必须幂等,如果返回错误，会一直重试;所以对服务要求很高，不能陷入死循环
	compensatePayFunc := func(ctx context.Context, uid int64, orderId string, value float64) error {
		_, err := wClient.CompensatePayment(ctx, &wallet.CompensatePaymentRequest{
			Uid:     uid,
			OrderId: orderId,
			Value:   value,
		})
		if err != nil {
			log.Println("CompensatePayment err:", err)
		}
		return err
	}

	setOrderFunc := func(ctx context.Context, orderId string, status int32) error {
		_, err := oClient.SetOrder(ctx, &order.SetOrderStatusRequest{
			OrderId: orderId,
			Status:  status,
		})
		if err != nil {
			log.Println("setOrderFunc err:", err)
		}
		return err
	}

	compensateSetOrderFunc := func(ctx context.Context, orderId string, status int32) error {
		_, err := oClient.SetOrder(ctx, &order.SetOrderStatusRequest{
			OrderId: orderId,
			Status:  status - 1,
		})
		if err != nil {
			log.Println("compensateSetOrderFunc err:", err)
		}
		return err
	}
	logPrefix := "saga_"
	store, err := kafka.NewKafkaStorage(zkAddrs, brokerAddrs, partitions, replicas, returnDuration, logPrefix,
		log.New(os.Stdout, "saga_", log.LstdFlags))
	if err != nil {
		return err
	}
	sec := saga.NewSEC(store, logPrefix)
	sec.AddSubTxDef("钱包支付", payFunc, compensatePayFunc)
	sec.AddSubTxDef("修改订单状态", setOrderFunc, compensateSetOrderFunc)

	// sagaID 要求全局不冲突，在事务未结束（正常、异常）时不能重复使用sagaID
	//var sagaID uint64 = 2

	// ctx贯穿所有action的执行,此时间不能太短（保证补偿能正确执行）
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = sec.StartSaga(ctx, sagaID).
		ExecSub("钱包支付", uid, orderId, price).
		ExecSub("修改订单状态", orderId, int32(1)).EndSaga()
	// err是action执行的err
	if err != nil {
		log.Println(err)
	}
	return err
	// 如果saga内部日志操作失败会直接panic，所以对kafka要求很高
	// 如果使用内存存储日志，程序崩溃后无法恢复
}

func TestRecover(t *testing.T) {
	// 启动Coordinator，会从kafka将异常的事务最后一条日志打印出来,日志中包含执行参数，手工处理
	// 如果使用内存存储，则无法恢复，只能从panic日志跟踪事务执行轨迹
	logPrefix := "saga_"
	store, err := kafka.NewKafkaStorage(zkAddrs, brokerAddrs, partitions, replicas, returnDuration, logPrefix,
		log.New(os.Stdout, "saga_", log.LstdFlags))
	if err != nil {
		t.Fatal(err)
	}
	sec := saga.NewSEC(store, logPrefix)
	err = sec.StartCoordinator()
	if err != nil {
		t.Log(err)
	}
}
