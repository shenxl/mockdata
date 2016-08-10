package models

type CompanyInstall struct {
	Id      int64 `json:"id"`
	Company int64 `json:"company"`
	Server  int64 `json:"server"`
	Sum     int64 `json:"sum"`
}

func (CompanyInstall) TableName() string {
	return "company_install"
}
