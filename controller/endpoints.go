package controller

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"github.com/sanyogpatel-tecblic/To-Do/model"
// )

// var DB *sql.DB

// func GetAllTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var tasks []model.Task
// 	rows, err := DB.Query("SELECT id,tasks FROM todo")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var task model.Task
// 		err := rows.Scan(&task.ID, &task.Tasks)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	json.NewEncoder(w).Encode(tasks)
// }
// func GetDoneTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var tasks []model.Task
// 	rows, err := DB.Query("select tasks from todo where done=1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rows.Next() {
// 		var task model.Task
// 		err := rows.Scan(&task.Tasks)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	json.NewEncoder(w).Encode(tasks)
// }
// func GetTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var task model.Task
// 	row := DB.QueryRow("SELECT id,tasks FROM todo WHERE id=$1", id)
// 	err = row.Scan(&task.ID, &task.Tasks)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(task)
// }

// func CreateTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var task model.Task
// 	err := json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = DB.Exec("INSERT INTO todo (tasks) VALUES ($1)", task.Tasks)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(task)
// }

// func UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	param := mux.Vars(r)

// 	id, err := strconv.Atoi(param["id"])

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var task model.Task

// 	json.NewDecoder(r.Body).Decode(&task)

// 	_, err = DB.Exec("update todo set tasks=$1 where id=$2", task.Tasks, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(task)
// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var task model.Task
// 	_, err = DB.Exec("delete from todo where id=$1", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(task)
// }

// func MarkAsDone(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	res, err := DB.Exec("update todo set done=1 where id=$1", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(res)
// }
