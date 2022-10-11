package ctxdata

import (
	"context"
	"fmt"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) string {
	var identity string

	identity = fmt.Sprintf("%v", ctx.Value(CtxKeyJwtUserId))

	//if jsonIdentity, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
	//	identity = jsonIdentity.String()
	//	//if int64Uid, err := jsonIdentity.Int64(); err == nil {
	//	//	uid = int64Uid
	//	//} else {
	//	//	logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
	//	//}
	//
	//
	//
	//}
	return identity
}
