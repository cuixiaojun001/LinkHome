package common

import (
	"context"
	"github.com/cuixiaojun001/linkhome/modules/common/dao"
	"github.com/cuixiaojun001/linkhome/modules/common/model"
	"github.com/cuixiaojun001/linkhome/third_party/qiniu"
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
