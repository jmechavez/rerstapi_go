package domain

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmechavez/restapi_go/insurance/errs"
)

type ClientRepositoryDb struct {
	clientDb *sql.DB
}

func (r ClientRepositoryDb) JustName(fname string) (*Client, error) {
	findNameSql := `SELECT fname FROM clients WHERE fname = $1`

	rows := r.clientDb.QueryRow(findNameSql, fname)
	var c Client
	err := rows.Scan(
		&c.Fname,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Client not found")
		} else {
			log.Println("Error while querying client table: " + err.Error())
			return nil, errors.New("Unexpected database Error")
		}
	}
	return &c, nil
}

func (r ClientRepositoryDb) ByName(id string) (*Client, *errs.AppError) {
	findNameSql := `SELECT fname, lname, birthdate, id_card_1, id_no_1, id_card_2, id_no_2, 
                birthplace, contact_no, status, gender 
                FROM clients WHERE fname = $1`

	rows := r.clientDb.QueryRow(findNameSql, id)
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
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Client not found")
		} else {
			log.Println("Error while querying client table: " + err.Error())
			return nil, errs.NewUnxpectedError("Unexpected database Error")
		}
	}
	return &c, nil
}

func (r ClientRepositoryDb) FindAll(status string) ([]Client, *errs.AppError) {
	var rows *sql.Rows
	var err error

	// Dynamic SQL query based on status
	if status == "" {
		findAllSql := `SELECT fname, lname, birthdate, id_card_1, id_no_1, id_card_2, id_no_2, 
                       birthplace, contact_no, status, gender 
                       FROM clients`
		rows, err = r.clientDb.Query(findAllSql)
	} else {
		findAllSql := `SELECT fname, lname, birthdate, id_card_1, id_no_1, id_card_2, id_no_2, 
                       birthplace, contact_no, status, gender 
                       FROM clients WHERE status = $1`
		rows, err = r.clientDb.Query(findAllSql, status)
	}

	// Handle query execution error
	if err != nil {
		log.Printf("Error while querying clients table: %v\n", err)
		return nil, errs.NewUnxpectedError("Unexpected database server error")
	}
	defer rows.Close()

	// Parse query results
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
			log.Printf("Error while scanning clients table: %v\n", err)
			return nil, errs.NewUnxpectedError("Error parsing client data")
		}
		clients = append(clients, c)
	}

	// Ensure rows are iterated without errors
	if err = rows.Err(); err != nil {
		log.Printf("Error after scanning rows: %v\n", err)
		return nil, errs.NewUnxpectedError("Unexpected row iteration error")
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
