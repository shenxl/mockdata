package models

import "time"

// Company 描述公司信息
type Company struct {
	Model
	Name            string    `json:"name"`
	BuyNumber       int       `json:"buyNumber"`
	GroupId         uint      `gorm:"column:group_id" json:"groupid"`
	Region          string    `json:"region"`
	Industry        string    `json:"industry"`
	Province        string    `json:"province"`
	City            string    `json:"city"`
	Address         string    `json:"address"`
	LengthOfService int       `json:"lengthofService"`
	ServiceDate     time.Time `json:"serviceDate"`
}

// 定义company 表名
func (Company) TableName() string {
	return "company"
}
