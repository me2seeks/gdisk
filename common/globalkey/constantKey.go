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

//文件类型
var Postfix map[string][]string

func init() {
	Postfix = make(map[string][]string)
	Postfix["IMAGE"] = []string{"jpg", "jpeg", "png", "gif", "bmp", "webp"}
	Postfix["VIDEO"] = []string{"mp4", "avi", "flv", "wmv", "mkv", "mov", "rmvb", "mpg", "vob", "mts", "3gp", "m4v", "m2ts", "ts", "m3u8"}
	Postfix["AUDIO"] = []string{"mp3", "wav", "wma", "ogg", "flac", "aac", "ape", "m4a", "amr", "mid", "midi", "rmi"}
	Postfix["APP"] = []string{"apk", "ipa", "exe", "msi", "dmg", "pkg", "deb", "rpm", "jad", "jar", "sis", "sisx", "xap", "xpi", "xapk", "xar", "xar", "xz", "xz", "zip", "7z", "rar", "tar", "gz", "bz2", "cab", "arj", "z", "lz", "lzma", "lzo", "lzop", "lz4", "lz5", "lzs"}
	Postfix["BT"] = []string{"torrent"}
	Postfix["OTHER"] = []string{"rar", "xp3", "tjs"}
}
