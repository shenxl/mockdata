package models

type Server struct {
	Model
	Ip          string `json:"ip"`
	Mac         string `json:"mac"`
	Description string `json:"description"`
	CompanyId   int64  `json:"companyid"`
}

func (Server) TableName() string {
	return "server"
}
