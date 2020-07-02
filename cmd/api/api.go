package api

import (
	_ "github.com/axengine/go-saga/storage/memory"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"microtrans/config"
	"time"
)

var service micro.Service

func init() {
	cfg := config.DefaultConfig
	microService := micro.NewService(micro.Name("microtrans.test"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
		micro.Registry(etcdv3.NewRegistry(registry.Addrs(cfg.Service.Registry...))),
	)
	go microService.Run()
}
