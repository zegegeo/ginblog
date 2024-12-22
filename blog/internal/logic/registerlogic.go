package logic

import (
	"context"
	"time"

	"ginblog/blog/internal/svc"
	"ginblog/blog/internal/types"
	"ginblog/blog/internal/model/user_model"
	"ginblog/blog/internal/biz"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (err error) {
	//1.根据用户名查询是否已经存在，如果存在则返回注册错误
	userModel := user_model.NewUsersModel(l.svcCtx.Mysql)
	user, err := userModel.FindByUsername(l.ctx,req.Username)
	if err != nil {
		l.Logger.Error("查询用户失败", err)
		return biz.DBError
	}
	if user != nil {
		//用户已经注册
		return biz.AlreadyRegister
	}
	//2.如果不存在，插入用户数据，注册成功	
	_, err = userModel.Insert(l.ctx, &user_model.Users{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return biz.DBError
	}
	return
}
