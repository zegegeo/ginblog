package logic

import (
	"context"

	"goblog/blog/internal/svc"
	"goblog/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogListLogic {
	return &GetBlogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlogListLogic) GetBlogList(req *types.GetBlogListReq) (resp *types.BlogListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
