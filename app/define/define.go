package define

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = "xgAQYVEcmlX7f197tVu77ybzAD4h8NPK"
var TencentSecretID = "AKIDMaeYIqBKGzUmrjQxGLDvQgkbFfqnW1HO"
var CosBucket = "https://buffet-1306963147.cos.ap-shanghai.myqcloud.com"
var CosFolderName = "buffet"
var AvatarBaseUrl = CosBucket + "/" + CosFolderName + "/avatars/"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2000-01-01 00:00:01"

var TokenExpire = 60 * 60 * 24 * 3        // 3 days
var RefreshTokenExpire = 60 * 60 * 24 * 7 // 7 days

var UserRepositoryMaxSize = 1000 * 1024 * 1024  // 1GB
var PublicRepositoryMaxSize = 500 * 1024 * 1024 // 1GB
