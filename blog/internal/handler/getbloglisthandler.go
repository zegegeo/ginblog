package handler

import (
	"net/http"

	"ginblog/blog/internal/logic"
	"ginblog/blog/internal/svc"
	"ginblog/blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBlogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBlogListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetBlogListLogic(r.Context(), svcCtx)
		resp, err := l.GetBlogList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
