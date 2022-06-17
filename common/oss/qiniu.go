package oss

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
)

const AK = "UZKiLt7Q4jhJA6z2xGo4jSpa-DyUPY5AsyzJ3RtA"
const SK = "SoH_ie2zGIb_NegryWa5CyMYYUpwBs9GfPPljUqS"
const Bucket = "chinaskill"
const Url = "rc4dfplxe.hd-bkt.clouddn.com/"

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(key string, data []byte) (string, error) {

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
	newUrl := Url + key
	return newUrl, nil
}

//返回上传凭证
func UploadCertificate(max int64) string {
	putPolicy := storage.PutPolicy{
		Scope:            Bucket,
		CallbackURL:      "http://api.example.com/qiniu/oss/callback",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"path":"$(x:path)","uid":"$(x:uid)"}`,
		CallbackBodyType: "application/json",
		FsizeLimit:       max,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)
	putPolicy.Expires = 1800

	return upToken
}

//下载url
func DownloadUrl(hash string) string {
	mac := qbox.NewMac(AK, SK)
	privateDownloadUrl := storage.MakePrivateURLv2(mac, Url, hash, 30)
	return privateDownloadUrl
}
