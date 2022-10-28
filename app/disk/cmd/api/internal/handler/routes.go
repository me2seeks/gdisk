// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	File "cloud-disk/app/disk/cmd/api/internal/handler/File"
	Share "cloud-disk/app/disk/cmd/api/internal/handler/Share"
	"cloud-disk/app/disk/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/share/detail",
				Handler: Share.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share/popular/list",
				Handler: Share.PopularShareListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/statistics",
				Handler: Share.StatisticsHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/file/public/list",
				Handler: File.PublicFileListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/share/create",
				Handler: Share.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share/save",
				Handler: Share.SaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/user/list",
				Handler: Share.UserShareListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/file/public/save",
				Handler: File.PublicFileSaveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/repository/save",
				Handler: File.UserRepositorySaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/list",
				Handler: File.UserFileListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/rename",
				Handler: File.UserFileNameUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/folder/create",
				Handler: File.UserFolderCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/file",
				Handler: File.UserFileDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/file",
				Handler: File.UserFileMoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: File.FileUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload/prepare",
				Handler: File.FileUploadPrepareHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload/chunk",
				Handler: File.FileUploadChunkHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload/chunk/complete",
				Handler: File.FileUploadChunkCompleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
