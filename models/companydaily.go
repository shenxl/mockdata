package models

// CompanyDaily 描述企业日报活 信息
type CompanyDaily struct {
	//  Id值为主键
	Id          int64 `json:"id"`
	CompanyId   int64 `json:"companyid"`
	ServerId    int64 `json:"serverid"`
	Year        int   `json:"year"`
	Month       int   `json:"month"`
	Day         int   `json:"day"`
	ActivitySum int64 `json:"activity_sum"`
	InstallSum  int64 `json:"install_sum"`
}

// TableName 设置表名 company_daily
func (CompanyDaily) TableName() string {
	return "company_daily"
}
