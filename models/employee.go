package models

import (
	"context"
	"fmt"
	"gorestapi/config"
	"log"
	"time"
)

type Employee struct {
	ID        int       `json:"id"`
	IDNumber  string       `json:"id_number"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	table          = "employee"
	layoutDatetime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]Employee, error) {
	var employees []Employee
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT id, id_number, name, location, created_at, updated_at FROM %v ORDER BY id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	defer rowQuery.Close()

	for rowQuery.Next() {
		var employee Employee
		if err = rowQuery.Scan(&employee.ID,
			&employee.IDNumber,
			&employee.Name,
			&employee.Location,
			&employee.CreatedAt,
			&employee.UpdatedAt); err != nil {
			fmt.Printf("%s \n", err)
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func InsertEmployee(ctx context.Context, employee Employee) error {
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}

	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (id_number, name, location, created_at, updated_at)" +
		"VALUES('%v','%v','%v','%v','%v')", table,
		employee.IDNumber,
		employee.Name,
		employee.Location,
		time.Now().Format(layoutDatetime),
		time.Now().Format(layoutDatetime))
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

func UpdateEmployee(ctx context.Context, employee Employee) error {

	db, err := config.MYSQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v set id_number = '%s', name ='%s', location = '%s', updated_at = '%v' where id = %d",
		table,
		employee.IDNumber,
		employee.Name,
		employee.Location,
		time.Now().Format(layoutDatetime),
		employee.ID,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}