// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-disk/app/community/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/posts/list",
				Handler: PostsListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/detail",
				Handler: PostsDetailHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/posts/create",
				Handler: PostsCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/update",
				Handler: PostsUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/posts/delete",
				Handler: PostsDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/comment/create",
				Handler: PostsCommentCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/posts/comment/delete",
				Handler: PostsCommentDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/comment",
				Handler: PostsCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/feedback/create",
				Handler: PostsFeedbackCreateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
