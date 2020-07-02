package order

import (
	"context"
	"fmt"
	"github.com/go-xorm/xorm"
	"microtrans/model"
	"microtrans/proto/order"
)

type OrderService struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *OrderService {
	return &OrderService{
		engine: engine,
	}
}

func (srv *OrderService) CreateOrder(ctx context.Context, in *order.CreateOrderRequest, out *order.Response) error {
	_, err := srv.engine.Insert(&model.DWOrder{
		OrderId: in.OrderId,
		Uid:     uint64(in.Uid),
		Memo:    fmt.Sprintf("goods:%d prive:%v", in.GetGoods(), in.GetPrice()),
		Status:  0,
	})
	return err
}

func (srv *OrderService) SetOrder(ctx context.Context, in *order.SetOrderStatusRequest, out *order.Response) error {
	//if in.Status == 1 {
	//	time.Sleep(time.Minute * 1000)
	//}
	var order model.DWOrder
	_, err := srv.engine.Where("order_id=?", in.OrderId).Get(&order)
	if err != nil {
		return err
	}
	order.Status = in.Status
	_, err = srv.engine.ID(order.Id).Cols("status").Update(&order)
	return err
}
