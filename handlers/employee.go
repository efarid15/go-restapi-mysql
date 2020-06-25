package handlers

import (
	"context"
	"encoding/json"
	"gorestapi/models"
	"gorestapi/utils"
	"log"
	"net/http"
)

func Employee(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		employees, err := models.GetAll(ctx)

		if err != nil {
			println(err)
		}

		utils.ResponseJSON(w, employees, http.StatusOK)
		return
	}


	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Use content type application / json", http.StatusBadRequest)
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

		if err := models.InsertEmployee(ctx, emp); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		/*res := map[string]string{
			"result": "Create Employee Success",
		}*/

		utils.ResponseJSON(w, emp, http.StatusCreated)
		return
	}

	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Use content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var emp models.Employee

		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := models.UpdateEmployee(ctx, emp); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, emp, http.StatusCreated)
		return
	}

	http.Error(w, "No Permission", http.StatusMethodNotAllowed)
	return

}
