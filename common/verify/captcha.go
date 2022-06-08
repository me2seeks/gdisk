package verify

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

// Captcha 方便后期扩展
type Captcha struct{}

// 单例
var captchaInstance *Captcha

func Instance() *Captcha {
	if captchaInstance == nil {
		captchaInstance = &Captcha{}
	}
	return captchaInstance
}

// CreateImage 创建图片验证码
func (this *Captcha) CreateImage() string {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	return captchaId
}

// Reload 重载
func (this *Captcha) Reload(captchaId string) bool {
	return captcha.Reload(captchaId)
}

// Verify 验证
func (this *Captcha) Verify(captchaId, val string) bool {
	return captcha.VerifyString(captchaId, val)
}

// GetImageByte 获取图片二进制流
func (this *Captcha) GetImageByte(captchaId string) []byte {
	var content bytes.Buffer
	err := captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return content.Bytes()
}
