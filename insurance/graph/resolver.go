package graph

import (
	"context"
	"github.com/jmechavez/restapi_go/insurance/service"
)

type Resolver struct {
	ClientService service.ClientService
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *queryResolver) Clients(ctx context.Context) ([]*Client, error) {
	clients, err := r.ClientService.GetAllClient()
	if err != nil {
		return nil, err
	}

	var result []*Client
	for _, c := range clients {
		result = append(result, &Client{
			Fname:      c.Fname,
			Lname:      c.Lname,
			Birthdate:  c.Birthdate,
			IdCard1:    c.IdCard1,
			IdNo1:      c.IdNo1,
			IdCard2:    c.IdCard2,
			IdNo2:      c.IdNo2,
			BirthPlace: c.BirthPlace,
			ContactNo:  c.ContactNo,
			Status:     c.Status,
			Gender:     c.Gender,
		})
	}
	return result, nil
}

func (r *queryResolver) ClientByName(ctx context.Context, fname string) (*Client, error) {
	client, err := r.ClientService.FindName(fname)
	if err != nil {
		return nil, err
	}

	return &Client{
		Fname:      client.Fname,
		Lname:      client.Lname,
		Birthdate:  client.Birthdate,
		IdCard1:    client.IdCard1,
		IdNo1:      client.IdNo1,
		IdCard2:    client.IdCard2,
		IdNo2:      client.IdNo2,
		BirthPlace: client.BirthPlace,
		ContactNo:  client.ContactNo,
		Status:     client.Status,
		Gender:     client.Gender,
	}, nil
}
