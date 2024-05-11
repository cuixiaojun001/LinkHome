package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/utils"
	"github.com/cuixiaojun001/LinkHome/modules/common/dao"
	"github.com/cuixiaojun001/LinkHome/modules/common/model"
	houseDao "github.com/cuixiaojun001/LinkHome/modules/house/dao"
	orderDao "github.com/cuixiaojun001/LinkHome/modules/order/dao"
	userDao "github.com/cuixiaojun001/LinkHome/modules/user/dao"
	"github.com/cuixiaojun001/LinkHome/third_party/qiniu"
	"reflect"
	"regexp"
)

func AreaInfo() (*AreaList, error) {
	var areaList []model.Province
	provinceList, err := dao.GetAllProvince()
	if err != nil {
		return nil, err
	}
	for _, province := range provinceList {
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
	return &AreaList{AreaList: areaList}, nil
}

func UploadFile(_ context.Context, filename string, data []byte) *UploadFileDataItem {
	key, url := qiniu.Client.UploadFile(data)
	return &UploadFileDataItem{
		FileName: filename,
		FileKey:  key,
		FileUrl:  url,
	}
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
