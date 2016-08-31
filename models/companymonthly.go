package models

type CompanyMonthly struct {
	Id          int64 `json:"id"`
	CompanyId   int64 `json:"companyid"`
	ServerId    int64 `json:"serverid"`
	Year        int   `json:"year"`
	Month       int   `json:"month"`
	ActivitySum int64 `json:"activity_sum"`
	ActivityMax int64 `json:"activity_max"`
	ActivityAvg int64 `json:"activity_avg"`
	InstallSum  int64 `json:"install_sum"`
	InstallMax  int64 `json:"install_max"`
	InstallAvg  int64 `json:"install_avg"`
}

func (CompanyMonthly) TableName() string {
	return "company_monthly"
}
