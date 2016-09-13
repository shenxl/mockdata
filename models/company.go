package models

// Company 描述公司信息
type Company struct {
	Model
	Name      string `json:"name"`
	GroupId   uint   `gorm:"column:group_id" json:"groupid"`
	Region    string `json:"region"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	Industry  string `json:"industry"`
	Important int    `json:"important"`
}

// 定义 company 表名
func (Company) TableName() string {
	return "company"
}
