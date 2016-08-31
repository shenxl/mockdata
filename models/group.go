package models

// Company 描述公司信息
type Group struct {
	Model
	GroupName      string `json:"groupName"`
	PurchaseNumber int    `json:"purchaseNumber"`
	Industry       string `json:"industry"`
	Sales          string `json:"sales"`
}

// 定义company 表名
func (Group) TableName() string {
	return "company_group"
}
