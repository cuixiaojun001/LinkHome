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

type NewsListRequest struct {
	QueryParams struct {
		ID int `json:"id"` // 主键id
	} `json:"query_params"` // QueryParams 房源列表查询参数
	Offset    int      `json:"offset"`    // Offset 分页偏移量
	Limit     int      `json:"limit"`     // Limit 每页显示数量
	Orderings []string `json:"orderings"` // Orderings 排序字段
}

type NewsListResponse struct {
	Total      int            `json:"total"`       // Total 总数
	HasMore    bool           `json:"has_more"`    // HasMore 是否有下一页
	NextOffset int            `json:"next_offset"` // NextOffset offset下次起步
	DataList   []NewsListItem `json:"data_list"`   // DataList 用户列表
}

type NewsListItem struct {
	ID       int    `json:"id"`        // 主键ID
	Title    string `json:"title"`     // 标题
	Content  string `json:"content"`   // 内容
	CreateTs int64  `json:"create_ts"` // 创建时间
	Scene    string `json:"scene"`     // 公告场景 notice:公告 advertising:广告
	State    string `json:"state"`     // 公告状态 normal:正常 deleted:删除
}
