package logic

import (
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/common/ctxdata"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
	identity := ctxdata.GetUidFromCtx(l.ctx)

	//
	//{"@timestamp":"2022-10-19T18:58:54.360+08:00","caller":"handler/authhandler.go:103","content":"authorize failed: token contains an invalid number of segments\n=\u003e GET /user/detail HTTP/1.0\r\nHost: 127.0.0.1\r\nConnection: close\r\nAccept: application/json\r\nAccept-Encoding: gzip, deflate, br\r\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\r\nAuthorization: undefined\r\nConnection: close\r\nDnt: 1\r\nReferer: http://127.0.0.1/\r\nRemote-Host: 172.19.0.1\r\nSec-Ch-Ua: \"Chromium\";v=\"106\", \"Microsoft Edge\";v=\"106\", \"Not;A=Brand\";v=\"99\"\r\nSec-Ch-Ua-Mobile: ?0\r\nSec-Ch-Ua-Platform: \"Linux\"\r\nSec-Fetch-Dest: empty\r\nSec-Fetch-Mode: cors\r\nSec-Fetch-Site: same-origin\r\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.30\r\nX-Forwarded-For: 172.19.0.1\r\nX-Real-Ip: 172.19.0.1\r\n\r\n","level":"error"}
	//

	userInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Identity: identity,
	})
	if err != nil {
		return nil, err
	}

	//var userInfo types.User
	//_ = copier.Copy(&userInfoResp, userInfo)

	println(userInfoResp.User.CreatedAt)
	println(userInfoResp.User.Avatar)

	return &types.UserInfoResp{
		UserInfo: types.User{
			Identity: identity,
			Email:    userInfoResp.User.Email,
			Name:     userInfoResp.User.Name,
			Sex:      userInfoResp.User.Sex,
			Avatar:   userInfoResp.User.Avatar,
			Info:     userInfoResp.User.Info,
			Capacity: 1048576000,
			CreateAt: userInfoResp.User.CreatedAt,
		},
	}, nil
}
