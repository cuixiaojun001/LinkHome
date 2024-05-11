package orm

// IQuery 通用查询接口
type IQuery interface {
	// ExactMatch 设置where子句匹配条件
	ExactMatch(key string, value interface{}) IQuery
	// Where 获取where子句查询条件
	Where() map[string]interface{}
	// SetPagination 设置翻页信息
	SetPagination(page int, size int) IQuery
	// Pagination 获取翻页信息
	Pagination() (page int, size int)
	// SetOrder 设置排序条件
	SetOrder(key string, desc bool) IQuery
	// Orders 获取排序条件
	Orders() []*Order
	// SetIn 设置 in 查询条件
	SetIn(key string, values []interface{}) IQuery
	// SetInString 设置 string 类型数据 in 查询条件
	SetInString(key string, values []string) IQuery
	// Range 设置范围查询条件
	Range(key string, min interface{}, max interface{}) IQuery
}

// Query 基础查询，实现了 IQuery
type Query struct {
	where  map[string]interface{} // where 存储where查询条件，键为字段名，值为字段值
	page   int                    // Page 存储当前页码
	limit  int                    // Limit 存储每页的记录数
	orders []*Order               // Order list
}

// NewQuery 创建 Query 查询对象，默认不分页, 不排序
func NewQuery() *Query {
	return &Query{where: make(map[string]interface{})}
}

type Order struct {
	Key    string
	IsDesc bool
}

// NewOrder creates new sort
func NewOrder(key string, desc bool) *Order {
	return &Order{
		Key:    key,
		IsDesc: desc,
	}
}

var _ IQuery = (*Query)(nil)

func (q *Query) ExactMatch(key string, value interface{}) IQuery {
	q.where[key] = value
	return q
}

func (q *Query) Where() map[string]interface{} {
	return q.where
}

func (q *Query) SetPagination(page int, size int) IQuery {
	q.page = page
	q.limit = size
	return q
}

func (q *Query) Pagination() (int, int) {
	return q.page, q.limit
}

func (q *Query) SetOrder(key string, isDesc bool) IQuery {
	s := NewOrder(key, isDesc)
	q.orders = append(q.orders, s)
	return q
}

func (q *Query) Orders() []*Order {
	return q.orders
}

func (q *Query) SetIn(key string, values []interface{}) IQuery {
	q.where[key] = NewOrList(values)
	return q
}

func (q *Query) SetInString(key string, values []string) IQuery {
	v := make([]interface{}, 0, len(values))
	for _, value := range values {
		v = append(v, value)
	}
	return q.SetIn(key, v)
}

func (q *Query) Range(key string, min interface{}, max interface{}) IQuery {
	q.where[key] = NewRange(min, max)
	return q
}

func (q *Query) FuzzyMatch(key string, value string) IQuery {
	q.where[key] = NewFuzzyMatchValue(value)
	return q
}

// OrList query
type OrList struct {
	Values []interface{}
}

// NewOrList creates a new or list
func NewOrList(values []interface{}) *OrList {
	return &OrList{
		Values: values,
	}
}

// Range query
type Range struct {
	Min interface{}
	Max interface{}
}

// NewRange creates a new range
func NewRange(min, max interface{}) *Range {
	return &Range{
		Min: min,
		Max: max,
	}
}

// FuzzyMatchValue query
type FuzzyMatchValue struct {
	Value string
}

// NewFuzzyMatchValue creates a new fuzzy match
func NewFuzzyMatchValue(value string) *FuzzyMatchValue {
	return &FuzzyMatchValue{
		Value: value,
	}
}
