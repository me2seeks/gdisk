// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Sex      int64  `json:"sex"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
	Capacity int    `json:"capacity"`
	CreateAt string `json:"create_at"`
}

type RegisterCountReq struct {
}

type RegisterCountResp struct {
	Count int64 `json:"count"`
}

type RegisterReq struct {
	Name     string `json:"name""`
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type RegisterResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WXMiniAuthResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}
