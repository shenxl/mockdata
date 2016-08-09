package models

type Company struct {
	ID   int64
	Company string
  Sn string
  // region string `json:"region"`
  // industry string `json:"industry"`
}

func (Company) TableName() string {
  return "company_sn"
}
