package models

type Server struct {
	Model
	Ip         string `json:"ip"`
	Mac        string `json:"mac"`
	Desciption string `json:"desciption"`
	Company    int64  `json:"companyid"`
}

func (Server) TableName() string {
	return "server"
}
