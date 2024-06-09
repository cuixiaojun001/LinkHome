package user

type UserAuthStatus string

// 用户实名认证状态
const (
	unauthorized string = "unauthorized" //  未实名认证
	authorized   string = "authorized"   //  已实名认证
	auditing     string = "auditing"     //  审核中
	unapprove    string = "unapprove"    //  审核未通过
)

const (
	normal  string = "normal"
	deleted string = "deleted"
)
