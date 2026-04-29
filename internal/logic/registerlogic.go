package logic

import (
	"context"
	"fmt"
	"time"

	"user_rpc/internal/model"
	"user_rpc/internal/svc"
	"user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 处理用户注册请求
// 接收请求参数，生成唯一的用户ID，并将用户记录插入到数据库中
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.EmptyResp, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		UserId:   uint64(time.Now().UnixNano()),
		Username: in.Username,
		Password: in.Password,
	})

	if err != nil {
		fmt.Println("err ", err.Error())
		return nil, err
	}

	fmt.Println("result", result)

	return &user.EmptyResp{}, nil
}
