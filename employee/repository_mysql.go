package employee

import (
	"context"
	"fmt"
	"gorestapi/config"
	"gorestapi/models"
	"gorestapi/utils"
	"log"
	"net/http"
	"time"
)

const (
	table = "employee"
	layout_DateTime = "2020-06-23 16:41:00"
)

func GetEmployees(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		employees, err := GetAll(ctx)
		if err != nil {
			println(err)
		}
		utils.ResponseJSON(w, employees, http.StatusOK)
		return
	}
	http.Error(w, "Data Tidak Ditemukan", http.StatusNotFound)
	return
}

func GetAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Koneksi database gagal", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var employee models.Employee
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&employee.ID,
			&employee.IDNumber,
			&employee.Name,
			&employee.Location,
			&employee.CreatedAt,
			&employee.UpdatedAt); err != nil {
			return nil, err
		}
		// change format date createdAt and updatedAt from string to datetime
		employee.CreatedAt, err = time.Parse(layout_DateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		employee.UpdatedAt, err = time.Parse(layout_DateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		// end change
		employees = append(employees, employee)
	}

	return employees, nil
}
