package config

var DefaultConfig = &Config{
	Mysql: Mysql{
		HostPort: "mysql1.dev.crycx.com:3306",
		Username: "root",
		DBName:   "microtrans",
		Password: "111111",
		MaxConns: 5,
		MaxIdle:  3,
	},
	Service: Service{
		OrderSrvName:  "microtrans.order",
		WalletSrvName: "microtrans.wallet",
		Registry:      []string{"http://etcd1.dev.crycx.com:2379", "http://etcd2.dev.crycx.com:2379", "http://etcd3.dev.crycx.com:2379"},
	},
}

type Config struct {
	Mysql   Mysql
	Service Service
}

type Mysql struct {
	HostPort string
	Username string
	DBName   string
	Password string
	MaxConns int
	MaxIdle  int
}

type Service struct {
	OrderSrvName  string
	WalletSrvName string
	Registry      []string
	SrvGather     string
	SrvApp        string
}
