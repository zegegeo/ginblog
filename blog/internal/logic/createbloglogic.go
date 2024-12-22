package logic

import (
	"context"

	"ginblog/blog/internal/svc"
	"ginblog/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBlogLogic) CreateBlog(req *types.CreateBlogReq) (resp *types.CreateBlogRes, err error) {
	// todo: add your logic here and delete this line

	return
}
