package logic

import (
	"context"

	"user_rpc/internal/svc"
	"user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResp{
		Token:         "token",
		Expire:        0,
		RefreshToken:  "refresh_token",
		RefreshExpire: 0,
	}, nil
}
