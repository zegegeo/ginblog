// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"GoBlog/api/internal/svc"
	user_handler "GoBlog/api/internal/handler/user"
	article_handler "GoBlog/api/internal/handler/article"
	comment_handler "GoBlog/api/internal/handler/comment"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: user_handler.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: user_handler.RegisterHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.DebugMiddleware, serverCtx.AuthMiddleware, serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/articles",
					Handler: article_handler.CreateArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/articles",
					Handler: article_handler.GetArticleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/articles/:article_id/comments",
					Handler: comment_handler.CreateCommentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/articles/:article_id/comments",
					Handler: comment_handler.GetCommentListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/articles/:id",
					Handler: article_handler.GetArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/articles/:id",
					Handler: article_handler.UpdateArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/articles/:id",
					Handler: article_handler.DeleteArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/comments/:id",
					Handler: comment_handler.DeleteCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)
}
