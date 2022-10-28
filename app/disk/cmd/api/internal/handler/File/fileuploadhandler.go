package File

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/oss"
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"

	"cloud-disk/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		// 生成文件hash, 判断文件是否已存在
		bt := make([]byte, fileHeader.Size)
		_, err = file.Read(bt)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(bt))

		rp := new(model.RepositoryPool)
		svcCtx.Engine.
			Where("hash = ?", hash).
			First(rp)
		// if err != nil {
		// 	httpx.OkJson(w, &types.FileUploadReply{
		// 		Msg: "error",
		// 	})
		// 	return
		// }
		// 文件已存在

		if rp.Id == 0 {
			cosPath, err := oss.CosUpload(r)
			if err != nil {
				return
			}
			req.Path = cosPath
		} else {
			req.Path = rp.Path
		}

		// to logic
		req.Name = fileHeader.Filename
		req.Hash = hash
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		if req.Ext == "" {
			req.Ext = ".unknown"
		}

		l := File.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		result.HttpResult(r, w, resp, err)
	}
}
