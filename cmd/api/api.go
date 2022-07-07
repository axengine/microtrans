package api

import (
	_ "github.com/axengine/go-saga/storage/memory"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"microtrans/config"
	"time"
)

var service micro.Service

func init() {
	cfg := config.DefaultConfig
	microService := micro.NewService(micro.Name("microtrans.test"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
		micro.Registry(etcd.NewRegistry(registry.Addrs(cfg.Service.Registry...))),
	)
	go microService.Run()
}
