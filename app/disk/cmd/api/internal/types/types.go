// Code generated by goctl. DO NOT EDIT.
package types

type File struct {
	Path     string `json:"path",default:"/"`
	FileName string `json:"fileName"`
	Size     int64    `json:"size"`
}

type UploadCertificateReq struct {
	Files []File `json:"files"`
}

type UploadCertificateResp struct {
	Certificate string `json:"certificate""`
}

type DownloadCertificateReq struct {
}

type DownloadCertificateResp struct {
	Certificate string `json:"certificate""`
}

type CallbackReq struct {
}

type CallbackResp struct {
}

type Folder struct {
	Id         int64 `json:"id"`
	UpdateTime int64 `json:"update_time"`
}

type FolderFile struct {
	Id         int64  `json:"id"`
	Size       int64  `json:"size"`
	Format     string `json:"format"`
	UpdateTime int64  `json:"update_time"`
	CreateTime int64  `json:"create_time"`
}

type CreateFolderReq struct {
	ParentFolderId string `json:"parentFolderId"`
	FolderName     string `json:"folderName"`
}

type CreateFolderResp struct {
	Id int64 `json:"id"`
}

type ListReq struct {
	Path string `json:"path"`
}

type ListResp struct {
	FolderList []Folder     `json:"folderList"`
	FileList   []FolderFile `json:"fileList"`
}

type MovedReq struct {
	Id       int64  `json:"id"`
	IsFile   bool   `json:"isFile"`
	SrcPath  string `json:"srcPath"`
	DestPath string `json:"destPath"`
}

type MovedResp struct {
}

type RenamedReq struct {
	Id      int64  `json:"id"`
	IsFile  bool   `json:"isFile"`
	NewName string `json:"name"`
}

type RenamedResp struct {
}

type StoreDetailReq struct {
}

type StoreDetailResp struct {
	CurrentSize int64 `json:"current_size"`
	MaxSize     int64 `json:"max_size"`
}
