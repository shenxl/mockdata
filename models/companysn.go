package models

type CompanySn struct {
	Model
	CompanyId uint   `json:"companyId"`
	Sn        string `json:"sn"`
}

func (CompanySn) TableName() string {
	return "company_sn"
}
