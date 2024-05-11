package orm

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SetQuery 设置 db 的查询条件
func SetQuery(orm *gorm.DB, query IQuery) *gorm.DB {
	if query == nil {
		return orm
	}

	page, size := query.Pagination()
	offset := (page - 1) * size

	orm = setQuery(orm, query)

	for _, order := range query.Orders() {
		if order.IsDesc {
			orm = orm.Order(order.Key + " DESC")
		} else {
			orm = orm.Order(order.Key + " ASC")
		}
	}
	// 需要链式调用，否则不会作用在一个查询上
	return orm.Offset(offset).Limit(size)
}

// setQuery 复合查询条件
func setQuery(orm *gorm.DB, query IQuery) *gorm.DB {
	where := make([]string, 0)
	args := make([]interface{}, 0)
	for k, v := range query.Where() {
		switch t := v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string: // 基础数据类型
			where = append(where, fmt.Sprintf("%s = ?", k))
			args = append(args, t)
		case *OrList:
			where = append(where, fmt.Sprintf("%s IN (?)", k))
			args = append(args, t.Values)
		case *Range:
			where = append(where, fmt.Sprintf("%s BETWEEN ? AND ?", k))
			args = append(args, t.Min, t.Max)
		case *FuzzyMatchValue:
			where = append(where, fmt.Sprintf("%s LIKE ?", k))
			args = append(args, "%"+t.Value+"%")
		default:
			// TODO 复杂数据类型
		}
	}
	if len(where) > 0 {
		qs := strings.Join(where, " AND ")
		orm = orm.Where(qs, args...)
	}
	return orm
}
