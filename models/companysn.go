package models

type CompanySn struct {
	Model
	Company int64 `json:"company"`
	Sn      int64 `json:"sn"`
}

func (CompanySn) TableName() string {
	return "company_sn"
}
