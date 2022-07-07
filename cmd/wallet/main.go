package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"microtrans/config"
	"microtrans/model"
	proto "microtrans/proto/wallet"
	"microtrans/service/wallet"
	"time"
	"xorm.io/core"
)

func main() {
	cfg := config.DefaultConfig

	engine := initEngine(cfg)

	var service micro.Service
	service = micro.NewService(
		micro.Name(cfg.Service.WalletSrvName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(etcd.NewRegistry(registry.Addrs(cfg.Service.Registry...))),
	)
	service.Init()

	walletSrv := wallet.New(engine)
	_ = proto.RegisterWalletHandler(service.Server(), walletSrv)

	if err := service.Run(); err != nil {
		return
	}
}

func initEngine(cfg *config.Config) *xorm.Engine {
	dial := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", cfg.Mysql.Username,
		cfg.Mysql.Password, cfg.Mysql.HostPort, cfg.Mysql.DBName)
	engine, err := xorm.NewEngine("mysql", dial)
	if err != nil {
		panic(err)
	}
	engine.SetMaxOpenConns(cfg.Mysql.MaxConns)
	engine.SetMaxIdleConns(cfg.Mysql.MaxIdle)
	engine.ShowSQL(true)

	//engine.SetLogger(xorm.NewSimpleLogger(os.Stdout))
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(model.DWAccountFlow), new(model.DWOrder), new(model.DWWallet))

	return engine
}
