package wallet

import (
	"context"
	"errors"
	"github.com/go-xorm/xorm"
	"microtrans/model"
	"microtrans/proto/wallet"
)

type WalletService struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *WalletService {
	return &WalletService{
		engine: engine,
	}
}

func (srv *WalletService) Payment(ctx context.Context, in *wallet.PaymentRequest, out *wallet.Response) error {
	var wallet model.DWWallet
	b, err := srv.engine.Where("uid=?", in.Uid).Get(&wallet)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("not found wallet")
	}

	var flow model.DWAccountFlow
	b, err = srv.engine.Where("out_order_id=?", in.OrderId).Get(&flow)
	if err != nil {
		return err
	}

	if b {
		return nil
	}

	sess := srv.engine.NewSession()
	sess.Begin()

	_, err = sess.ID(wallet.Id).And("balance=?", wallet.Balance).Cols("balance").Update(&model.DWWallet{
		Balance: wallet.Balance - in.Value,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	_, err = sess.Insert(&model.DWAccountFlow{
		UID:        uint64(in.Uid),
		Amount:     in.Value,
		Balance:    wallet.Balance - in.Value,
		Direction:  "-",
		OutOrderId: in.OrderId,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	sess.Commit()
	return nil
}

func (srv *WalletService) CompensatePayment(ctx context.Context, in *wallet.CompensatePaymentRequest, out *wallet.Response) error {
	var wallet model.DWWallet
	b, err := srv.engine.Where("uid=?", in.Uid).Get(&wallet)
	if err != nil {
		return err
	}
	if !b {
		//return errors.New("not found wallet")
		return nil
	}

	var flow model.DWAccountFlow
	b, err = srv.engine.Where("out_order_id=?", in.OrderId).And("direction=?", "-").Get(&flow)
	if err != nil {
		return err
	}

	if !b {
		//return errors.New("not found flow log")
		return nil
	}

	// 补偿过了 不用再补偿
	b, err = srv.engine.Where("out_order_id=?", in.OrderId).And("direction=?", "+").Get(&flow)
	if err != nil {
		return err
	}

	if b {
		return nil
	}

	// 开始补偿
	sess := srv.engine.NewSession()
	sess.Begin()

	_, err = sess.ID(wallet.Id).And("balance=?", wallet.Balance).Cols("balance").Update(&model.DWWallet{
		Balance: wallet.Balance + in.Value,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	_, err = sess.Insert(&model.DWAccountFlow{
		UID:        uint64(in.Uid),
		Amount:     in.Value,
		Balance:    wallet.Balance + in.Value,
		Direction:  "+",
		OutOrderId: in.OrderId,
	})
	if err != nil {
		sess.Rollback()
		return err
	}

	sess.Commit()

	return nil
}
