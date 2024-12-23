package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"goblog/blog/internal/biz"
	"goblog/blog/internal/model/user_model"
	"goblog/blog/internal/svc"
	"goblog/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil,biz.TokenError
	}
	userModel := user_model.NewUsersModel(l.svcCtx.Mysql)
	user, err := userModel.FindOne(l.ctx, userId)
	if err != nil &&(errors.Is(err,user_model.ErrNotFound)||errors.Is(err,sql.ErrNoRows)){
		return nil,biz.TokenError
	}
	if err != nil {
		l.Logger.Error("查询用户信息失败", err)
		return nil,biz.DBError
	}
	return &types.User{
		ID:       user.Id,
		Username: user.Username,
	}, nil
}
