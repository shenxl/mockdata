package models

type CompanyMonthly struct {
	Id          int64 `json:"id"`
	Company     int64 `json:"company"`
	Server      int64 `json:"server"`
	Year        int64 `json:"year"`
	Month       int64 `json:"month"`
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
