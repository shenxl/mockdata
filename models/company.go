package models

type Company struct {
	Id   int64  `json:"id"`
	Company string `json:"company"`
  Group string `json:"group"`
  Region string `json:"region"`
  Industry string `json:"industry"`
}

func (Company) TableName() string {
  return "company"
}
