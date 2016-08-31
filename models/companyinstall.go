package models

type CompanyInstall struct {
	Id        int64 `json:"id"`
	CompanyId int64 `json:"companyid"`
	ServerId  int64 `json:"serverid"`
	Sum       int64 `json:"sum"`
}

func (CompanyInstall) TableName() string {
	return "company_install"
}
