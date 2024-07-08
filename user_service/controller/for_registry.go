package controller

import (
	"context"
	"user_service/helper"
	"user_service/pb/user_registry"
)

func (c *UserRegistry) GetBalance(ctx context.Context, in *user_registry.BalanceReq) (*user_registry.BalanceResp, error) {
	resp, err := c.UR.GetInfo(uint(in.UserId))
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return &user_registry.BalanceResp{Deposit: resp.Deposit}, nil
}

func (c *UserRegistry) UpdateBalance(ctx context.Context, in *user_registry.BalanceUpdate) (*user_registry.UpdateResp, error) {
	err := c.UR.SubstractBalance(uint(in.UserId), in.NewBalance)
	if err != nil {
		return nil, helper.ParseErrorGRPC(err)
	}

	return nil, nil
}
