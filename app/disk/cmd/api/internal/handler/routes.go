// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "cloud-disk/app/disk/cmd/api/internal/handler/file"
	fileFolder "cloud-disk/app/disk/cmd/api/internal/handler/fileFolder"
	store "cloud-disk/app/disk/cmd/api/internal/handler/store"
	"cloud-disk/app/disk/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/certificate/upload",
				Handler: file.UploadCertificateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/certificate/download",
				Handler: file.DownloadCertificateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/certificate/callback",
				Handler: file.CallbackHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/disk/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/folder",
				Handler: fileFolder.FoldercreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/folder/list",
				Handler: fileFolder.FolderlistHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/folder/move",
				Handler: fileFolder.MoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/folder/rename",
				Handler: fileFolder.RenameHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/disk/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/store",
				Handler: store.StoreHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/disk/v1"),
	)
}