package models

type Company struct {
	id   int64  `json:"id"`
	company string `json:"company"`
  sn string `json:"sn"`
  // region string `json:"region"`
  // industry string `json:"industry"`
}

func (Company) TableName() string {
  return "company_sn"
}
