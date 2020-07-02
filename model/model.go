package model

import "time"

type DWWallet struct {
	Id        int64     `xorm:"not null pk autoincr BIGINT(20)"`
	UID       uint64    `xorm:"not null unique(u_wallet) BIGINT(20) comment('用户ID')"`
	Balance   float64   `xorm:"not null DECIMAL(40,20) comment('余额')"`
	Status    int32     `xorm:"DEFAULT 0 TINYINT(4) comment('钱包状态 0 正常')"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type DWAccountFlow struct {
	Id         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	UID        uint64    `xorm:"not null index(idx_phone_coin) BIGINT(20) comment('用户ID')"`
	Amount     float64   `xorm:"not null DECIMAL(40,20) comment('变动金额')"`
	Balance    float64   `xorm:"not null DECIMAL(40,20) comment('变动后余额')"`
	Direction  string    `xorm:"not null VARCHAR(2) comment('变动方向:+,-')"`
	OutOrderId string    `xorm:"not null index VARCHAR(64) comment('关联的业务Id，比如订单id，交易hash等')"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type DWOrder struct {
	Id        int64 `xorm:"not null pk autoincr BIGINT(20)"`
	OrderId   string
	Uid       uint64
	Memo      string
	Status    int32     `xorm:"DEFAULT 0 TINYINT(4) comment('订单状态 0-待支付 1-已支付 2-支付失败')"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
