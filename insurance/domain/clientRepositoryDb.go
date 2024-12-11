package domain

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ClientRepositoryDb struct {
	clientDb *sql.DB
}

func (r ClientRepositoryDb) FindName() ([]Client, error) {
	findNameSql := "Select fname, lname"
	rows, err := r.clientDb.Query(findNameSql)
	if err != nil {
		log.Println("Error while querying client table: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	clients := make([]Client, 0)
	for rows.Next() {
		var c Client
		err := rows.Scan(
			&c.Fname,
			&c.Lname,
		)
		if err != nil {
			log.Println("Error while scanning client table: " + err.Error())
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func (r ClientRepositoryDb) FindAll() ([]Client, error) {
	findAllSql := "SELECT fname, lname, birthdate, id_card_1, id_no_1, id_card_2, id_no_2, birthplace, contact_no, status, gender FROM clients"
	rows, err := r.clientDb.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying client table: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	clients := make([]Client, 0)
	for rows.Next() {
		var c Client
		err := rows.Scan(
			&c.Fname,
			&c.Lname,
			&c.Birthdate,
			&c.IdCard1,
			&c.IdNo1,
			&c.IdCard2,
			&c.IdNo2,
			&c.BirthPlace,
			&c.ContactNo,
			&c.Status,
			&c.Gender,
		)
		if err != nil {
			log.Println("Error while scanning client table: " + err.Error())
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func NewClientRepositoryDb() ClientRepositoryDb {
	connStr := "user=admin password=admin123 dbname=insurance sslmode=disable" // Adjust sslmode if needed
	clientDb, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return ClientRepositoryDb{clientDb}
}
