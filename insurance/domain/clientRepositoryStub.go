package domain

type ClientRepositoryStub struct {
	client []Client
}

func (r ClientRepositoryStub) FindAll() ([]Client, error) {
	return r.client, nil
}

func NewClientRepositoryStub() ClientRepositoryStub {
	clients := []Client{
		{
			"John Michael",
			"Echavez",
			"09-16-1990",
			"National Id",
			"001-001-001",
			"Passport",
			"002-002-002",
			"General Santos City",
			"0900000001",
			1,
			"Male",
		},
		{
			"Jm",
			"Echavez",
			"09-16-1990",
			"National Id",
			"003-003-003",
			"Passport",
			"004-004-004",
			"General Santos City",
			"0900000002",
			1,
			"Male",
		},
	}
	return ClientRepositoryStub{clients}
}
