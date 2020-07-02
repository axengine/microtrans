package api

import (
	"github.com/axengine/go-saga"
	_ "github.com/axengine/go-saga/storage/kafka"
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

	initKafka()
}

func initKafka() {
	saga.StorageConfig.Kafka.ZkAddrs = []string{"192.168.10.32:2181"}
	saga.StorageConfig.Kafka.BrokerAddrs = []string{"192.168.10.32:9092"}
	saga.StorageConfig.Kafka.Partitions = 1
	saga.StorageConfig.Kafka.Replicas = 1
	saga.StorageConfig.Kafka.ReturnDuration = 50 * time.Millisecond
}
