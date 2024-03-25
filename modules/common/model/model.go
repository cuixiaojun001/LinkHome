package model

type District struct {
	ID       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	ParentID int    `gorm:"parent_id" json:"parent_id"`
}

func (c *District) TableName() string {
	return "areas"
}

type City struct {
	ID           int        `gorm:"id" json:"id"`
	Name         string     `gorm:"name" json:"name"`
	ParentID     int        `gorm:"parent_id" json:"parent_id"`
	DistrictList []District `gorm:"-" json:"district_list"`
}

func (c *City) TableName() string {
	return "areas"
}

type Province struct {
	ID       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	ParentId int    `gorm:"parent_id" json:"-"`
	CityList []City `gorm:"-" json:"city_list"`
}

func (c *Province) TableName() string {
	return "areas"
}
