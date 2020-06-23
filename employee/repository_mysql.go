package employee

import (
	"context"
	"encoding/json"
	"fmt"
	"gorestapi/config"
	"gorestapi/models"
	"gorestapi/utils"
	"log"
	"net/http"
	"time"
)

const (
	table          = "employee"
	layoutDatetime = "2006-01-02 15:04:05"
)

func Employees(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		employees, err := GetAll(ctx)

		if err != nil {
			println(err)
		}
		//var dataEmployee interface{} = employees

		utils.ResponseJSON(w, employees, http.StatusOK)
		return
	}

	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var emp models.Employee

		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		if err := InsertEmployee(ctx, emp); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"result": "Create Employee Success",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return

	}
	http.Error(w, "No Permission", http.StatusMethodNotAllowed)
	return

}

func InsertEmployee(ctx context.Context, employee models.Employee) error {
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}

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


func GetAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Koneksi database gagal", err)
	}

	queryText := fmt.Sprintf("SELECT id, id_number, name, location, created_at, updated_at FROM %v ORDER BY id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {

		var employee models.Employee


		if err = rowQuery.Scan(&employee.ID,
			&employee.IDNumber,
			&employee.Name,
			&employee.Location,
			&employee.CreatedAt,
			&employee.UpdatedAt); err != nil {
			fmt.Printf("%s \n", err)
			return nil, err
		}
		// change format date createdAt and updatedAt from string to datetime

        employees = append(employees, employee)

	}

	return employees, nil
}
