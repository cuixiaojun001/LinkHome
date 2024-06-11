package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cuixiaojun001/LinkHome/common/cache"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"reflect"
	"regexp"
	"sync"

	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/utils"
	"github.com/cuixiaojun001/LinkHome/modules/common/dao"
	"github.com/cuixiaojun001/LinkHome/modules/common/model"
	houseDao "github.com/cuixiaojun001/LinkHome/modules/house/dao"
	orderDao "github.com/cuixiaojun001/LinkHome/modules/order/dao"
	userDao "github.com/cuixiaojun001/LinkHome/modules/user/dao"
	"github.com/cuixiaojun001/LinkHome/third_party/qiniu"
)

type ICommonService interface {
	// AreaInfo 获取省市区信息
	AreaInfo() (*AreaList, error)
	// UploadFile 上传文件
	UploadFile(_ context.Context, filename string, data []byte) *UploadFileDataItem
	// GetNews 获取公告资讯列表
	GetNews(_ context.Context, req *NewsListRequest) (*NewsListResponse, error)
}

type CommonService struct {
	cache cache.Cache
}

var once sync.Once
var commonManager ICommonService

func GetCommonManager() ICommonService {
	once.Do(func() {
		commonManager = &CommonService{
			cache: cache.New("area"),
		}
	})
	return commonManager
}

var _ ICommonService = (*CommonService)(nil)

func (s *CommonService) AreaInfo() (*AreaList, error) {
	result := &AreaList{}
	if exist, _ := s.cache.Get(context.TODO(), "info", &result); exist {
		return result, nil
	}

	var areaList []model.Province
	provinceList, err := dao.GetAllProvince()
	logger.Debugw("provinceList", "provinceList", provinceList)
	if err != nil {
		return nil, err
	}
	for _, province := range provinceList {
		if province.Name == "香港特别行政区" {
			logger.Debugw("province", "province", province)
		}
		areaItem := model.Province{
			ID:   province.ID,
			Name: province.Name,
		}

		cityList, err := dao.GetCityByID(province.ID)
		if err != nil {
			return nil, err
		}

		for i, city := range cityList {
			districtList, err := dao.GetDistrictByID(city.ID)
			if err != nil {
				return nil, err
			}
			cityList[i].DistrictList = districtList
		}
		areaItem.CityList = cityList
		areaList = append(areaList, areaItem)
	}
	result.AreaList = areaList
	if err := s.cache.Set(context.TODO(), "info", result); err != nil {
		logger.Errorw("cache set failed", "err", err)
	}

	return &AreaList{AreaList: areaList}, nil
}

func (s *CommonService) UploadFile(_ context.Context, filename string, data []byte) *UploadFileDataItem {
	key, url := qiniu.Client.UploadFile(data)
	return &UploadFileDataItem{
		FileName: filename,
		FileKey:  key,
		FileUrl:  url,
	}
}

func (s *CommonService) GetNews(_ context.Context, req *NewsListRequest) (*NewsListResponse, error) {
	filter := orm.NewQuery()
	if req.QueryParams.ID != 0 {
		filter.ExactMatch("id", req.QueryParams.ID)
	}
	if req.Limit != 0 && req.Offset != 0 {
		filter.SetPagination(req.Offset, req.Limit)
	}

	noticeList, err := dao.GetSystemNoticeList(filter)
	if err != nil {
		logger.Errorw("GetSystemNoticeList failed", "err", err)
		return nil, err
	}

	var newsList []NewsListItem
	for _, notice := range noticeList {
		newsItem := NewsListItem{
			ID:       notice.ID,
			Title:    notice.Title,
			Content:  notice.Content,
			CreateTs: notice.CreatedAt.Unix(),
		}
		newsList = append(newsList, newsItem)
	}

	return &NewsListResponse{
		DataList:   newsList,
		Total:      len(newsList),
		HasMore:    false,
		NextOffset: 0,
	}, nil
}

// GenerateContractContent 生成电子合同内容
// orderID 订单id
// templateID 电子合同模板id 默认1
func GenerateContractContent(_ context.Context, orderID int, templateID int) (string, error) {
	logger.Debugw("GenerateContractContent", "orderID", orderID, "templateID", templateID)
	templateID = 1
	order, err := orderDao.GetOrder(orderID)
	if err != nil {
		return "", err
	}
	templateContract, err := dao.GetContractTemplate(templateID)
	if err != nil {
		return "", err
	}

	var renderParams model.RenderParams
	err = json.Unmarshal(templateContract.RenderParams, &renderParams)
	if err != nil {
		return "", err
	}
	for renderKey, dataSrc := range renderParams {
		var modelID int
		var dbModel interface{}
		switch dataSrc.DBModelManager {
		case "OrderManager":
			modelID = order.ID
			dbModel, err = orderDao.GetOrder(modelID)
			if err != nil {
				return "", err
			}

		case "UserProfileManager":
			if dataSrc.Role == "tenant" {
				modelID = order.TenantID
			} else if dataSrc.Role == "landlord" {
				modelID = order.LandlordID
			}
			dbModel, err = userDao.GetUserProfile(modelID)
			if err != nil {
				return "", err
			}
		case "HouseInfoManager":
			modelID = order.HouseID
			dbModel, err = houseDao.GetHouseInfo(modelID)
			if err != nil {
				return "", err
			}

		default:
			return "", fmt.Errorf("unknown DBModelManager: %s", dataSrc.DBModelManager)
		}
		// dbModel是interface{}类型，需要反射
		val := reflect.ValueOf(dbModel)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		goFieldName := utils.ToCamelCase(dataSrc.FieldName)
		fieldVal := val.FieldByName(goFieldName)
		renderValue := fieldVal.String()

		re := regexp.MustCompile("{{\\s*" + renderKey + "\\s*}}")
		templateContract.TemplateContent = re.ReplaceAllString(templateContract.TemplateContent, renderValue)
	}

	return templateContract.TemplateContent, nil
}
