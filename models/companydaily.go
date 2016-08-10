package models

type CompanyDaily struct {
	Id          int64  `json:"id"`
	Company     int64  `json:"company"`
	Server      string `json:"server"`
	Year        int64  `json:"year"`
	Month       int64  `json:"month"`
	Day         int64  `json:"day"`
	ActivitySum int64  `json:"activity_sum"`
	InstallSum  int64  `json:"install_sum"`
}

func (CompanyDaily) TableName() string {
	return "company_daily"
}
