package handler

import (
	"encoding/json"
	"net/http"
	"services/api-service/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

var emp []model.Employee

func Init() { 
	emp = append(emp, model.Employee{
		ID: 1,
		FirstName: "Walter",
		SecondName: "White",
		LastName: "Hartwell",
		Grade: 5,
	})

	emp = append(emp, model.Employee{
		ID: 2,
		FirstName: "Jessy",
		SecondName: "Pinkman",
		LastName: "Fucking",
		Grade: 4,
	})

	emp = append(emp, model.Employee{
		ID: 3,
		FirstName: "Gus",
		SecondName: "Fring",
		LastName: "Fringovich",
		Grade: 3,
	})

	emp = append(emp, model.Employee{
		ID: 4,
		FirstName: "Soul",
		SecondName: "Goodman",
		LastName: "Hankovich",
		Grade: 2,
	})
}

func GetEmployees(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}


func GetEmployee(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil { 
		http.Error(w, "Неверный id", http.StatusBadRequest)
		return
	}

	for _, employee := range emp { 
		if employee.ID == id { 
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	http.Error(w, "Сотрудник не найден", http.StatusNotFound)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil { 
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return 
	}

	employee.ID = len(emp) + 1
	emp = append(emp, employee)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil { 
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return 
	}

	var updateEmp model.Employee
	err = json.NewDecoder(r.Body).Decode(&updateEmp)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return 
	}

	for i, employee := range emp { 
		if employee.ID == id { 
			updateEmp.ID = id 
			emp[i] = updateEmp
			json.NewEncoder(w).Encode(updateEmp)
			return 
		}
	}
	http.Error(w, "Сотрудник не найден", http.StatusNotFound)
}