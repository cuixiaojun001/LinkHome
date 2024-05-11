package common

import "github.com/cuixiaojun001/LinkHome/modules/common/model"

type AreaList struct {
	AreaList []model.Province `json:"area_list"`
}

type UploadFileDataItem struct {
	FileName string `json:"file_name"` // FileName 文件名
	FileKey  string `json:"file_key"`  // FileKey 文件key
	FileUrl  string `json:"file_url"`  // FileUrl 文件url
}
