package upload

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
)

const AK = "UZKiLt7Q4jhJA6z2xGo4jSpa-DyUPY5AsyzJ3RtA"
const SK = "SoH_ie2zGIb_NegryWa5CyMYYUpwBs9GfPPljUqS"

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(key string, data []byte) (string, error) {
	var Bucket = "chinaskill"
	var Url = "rc4dfplxe.hd-bkt.clouddn.com/"
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	Url = Url + key
	return Url, nil
}
