package models

// VCompanyMonthly 按公司统计报活情况
type VCompanyMonthly struct {
	Name         string `json:"name"`
	CompanyId    uint   `gorm:"column:company_id" json:"companyid"`
	ServerId     uint   `json:"serverid"`
	Year         int    `json:"year"`
	Month        int    `json:"month"`
	Date         string `json:"date"`
	ActivitySum  int    `json:"activity_sum"`
	InstallSum   int    `json:"install_sum"`
	ActivityMax  int    `json:"acitvity_max"`
	InstallMax   int    `json:"install_max"`
	ActivityAvg  int    `json:"acitvity_avg"`
	InstallAvg   int    `json:"install_avg"`
	InstallTotal int    `json:"install_total"`
	BuyTotal     int    `json:"buy_total"`
	Region       string `json:"region"`
	Type         string `json:"type"`
	Industry     string `json:"industry"`
	Important    int    `json:"important"`
}

// TableName 定义 VCompanyMonthly 表名
func (VCompanyMonthly) TableName() string {
	return "v_company_monthly"
}
