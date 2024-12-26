package domain

import "github.com/jmechavez/restapi_go/insurance/errs"

type Client struct {
	Fname      string `json:"fname"`
	Lname      string `json:"lname"`
	Birthdate  string `json:"birthdate"`
	IdCard1    string `json:"id_card_1"`
	IdNo1      string `json:"id_no_1"`
	IdCard2    string `json:"id_card_2"`
	IdNo2      string `json:"id_no_2"`
	BirthPlace string `json:"birthplace"`
	ContactNo  string `json:"contact_no"`
	Status     string `json:"status"`
	Gender     string `json:"gender"`
}

type ClientRepository interface {
	FindAll(status string) ([]Client, *errs.AppError)
	ByName(string) (*Client, *errs.AppError)
	JustName(string) (*Client, error)
}
