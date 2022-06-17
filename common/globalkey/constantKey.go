package globalkey

/**
global constant key
*/

//软删除
var DelStateNo int64 = 0  //未删除
var DelStateYes int64 = 1 //已删除

//时间格式化模版
var DateTimeFormatTplStandardDateTime = "Y-m-d H:i:s"
var DateTimeFormatTplStandardDate = "Y-m-d"
var DateTimeFormatTplStandardTime = "H:i:s"

//移动类型
var DeleteTypeMove int64 = -1 //移动
var RenameTypeMove int64 = -2
