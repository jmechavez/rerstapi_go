package domain

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
	Status     bool   `json:"status"`
	Gender     string `json:"gender"`
}

type ClientRepository interface {
	FindAll() ([]Client, error)
	ByName(string) (*Client, error)
	JustName(string) (*Client, error)
}
