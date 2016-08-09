package models

type Company struct {
	id   int64
	company string
  sn string
  // region string `json:"region"`
  // industry string `json:"industry"`
}

func (Company) TableName() string {
  return "company_sn"
}
