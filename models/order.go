package models

import "time"

// Order 描述订单信息
type Order struct {
	Model
	CompanyId          uint      `json:"company_id"`
	GroupId            uint      `json:"group_id"`
	Sales              string    `json:"sales"`
	Sns                string    `json:"sns"`
	OrderType          string    `json:"order_type"`
	OrderArea          string    `json:"order_area"`
	OrderName          string    `json:"order_name"`
	OrderNumber        int64     `json:"order_number"`
	LengthOfService    int       `json:"lengthofService"`
	ServiceDate        time.Time `json:"serviceDate"`
	AuthorizationYears int       `json:"authorization_years"`
	AuthorizationDate  time.Time `json:"authorization_date"`
}

// 定义 order 表名
func (Order) TableName() string {
	return "company_order"
}
