package models

type Company struct {
	Model
	Company  string `json:"company"`
	Group    string `json:"group"`
	Region   string `json:"region"`
	Industry string `json:"industry"`
}

func (Company) TableName() string {
	return "company"
}
