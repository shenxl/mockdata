package models

// Company 描述公司信息
type Group struct {
	Model
	GroupName string `json:"groupName"`
	Region    string `json:"region"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	Industry  string `json:"industry"`
	Important int    `json:"important"`
}

// 定义company 表名
func (Group) TableName() string {
	return "company_group"
}
