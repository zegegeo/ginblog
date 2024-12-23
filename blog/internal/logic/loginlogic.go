package logic

import (
	"context"
	"time"

	"goblog/blog/internal/biz"
	"goblog/blog/internal/svc"
	"goblog/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goblog/blog/internal/model/user_model"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	//1.校验用户名密码是否合法
	userModel := user_model.NewUsersModel(l.svcCtx.Mysql)
	user, err := userModel.FindByUsernameAndPassword(l.ctx,req.Username, req.Password)
	if err != nil {
		l.Logger.Error(err)
		return nil, biz.DBError
	}
	if user == nil {
		return nil, biz.UserNameAndPasswordError
	}
	//2.登陆成功，生成token
	secretkey := l.svcCtx.Config.Auth.SecretKey
	expire := l.svcCtx.Config.Auth.Expire
	token, err := biz.GetJwtToken(secretkey, time.Now().Unix(), expire, user.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, biz.TokenError
	}
	resp = &types.LoginRes{
		Token: token,
	}
	return
}
