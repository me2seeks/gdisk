package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

// 用户模块
const ErrUserPwdError uint32 = 200001
const ErrUserAlreadyRegisterError uint32 = 200002

// 验证
const VERIFY_ERROR uint32 = 600001

// 七牛云
const ErrorUploadFile uint32 = 700001
