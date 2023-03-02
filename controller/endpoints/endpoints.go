package endpoints

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sanyogpatel-tecblic/To-Do/controller/model"
)

type Response struct {
	Statuscode int         `json:"status"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func GenerateResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		Statuscode: statusCode,
		Message:    message,
		Data:       data,
	}
}
func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var task model.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("INSERT INTO todo (tasks) VALUES ($1)", task.Tasks)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(task)
	}
}
func GetAllTasks(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var tasks []model.Task
		// var row_var error

		rows, err := DB.Query("SELECT id,tasks FROM todo")
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var task model.Task
			row_var := rows.Scan(&task.ID, &task.Tasks)
			if row_var != nil {
				log.Fatal(row_var)
			}
			tasks = append(tasks, task)
		}
		json.NewEncoder(w).Encode(tasks)
		// defer rows.Close()
	}
}
func GetDoneTasks(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var tasks []model.Task
		rows, err := DB.Query("select id,tasks from todo where done=1")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var task model.Task
			err := rows.Scan(&task.ID, &task.Tasks)
			if err != nil {
				log.Fatal(err)

			}
			tasks = append(tasks, task)
		}
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
func GetTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}
		var task model.Task
		row := db.QueryRow("SELECT id,tasks FROM todo WHERE id=$1", id)
		err = row.Scan(&task.ID, &task.Tasks)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		param := mux.Vars(r)

		id, err := strconv.Atoi(param["id"])

		if err != nil {
			log.Fatal(err)
		}
		var task model.Task

		json.NewDecoder(r.Body).Decode(&task)

		_, err = db.Exec("update todo set tasks=$1 where id=$2", task.Tasks, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}
		var task model.Task
		_, err = db.Exec("delete from todo where id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(task)
	}
}

func MarkAsDone(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id := params["id"]
		res, err := db.Exec("update todo set done=1 where id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(res)
	}
}

// ---------------------------------------------------------------------------------//
// ------------------------------------Users---------------------------------------//
func RegisterUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO users (username,password) VALUES ($1,$2)", user.UserName, user.Password)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}
func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []model.User

		rows, err := db.Query("select * from users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var user model.User

			rowscn := rows.Scan(&user.ID, &user.UserName, &user.Password)
			if rowscn != nil {
				log.Fatal(rowscn)
			}
			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users)
	}
}
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user model.User
		param := mux.Vars(r)
		id, err := strconv.Atoi(param["id"])
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(r.Body).Decode(&user)
		update, _ := db.Exec("update users set username=$1,password=$2 where id=$3", user.UserName, user.Password, id)

		json.NewEncoder(w).Encode(update)
	}
}
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []model.User
		var user model.User
		if r.Body == nil {
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
			return
		}
		json.NewDecoder(r.Body).Decode(&user)

		rows, _ := db.Query("select username,password from users where username=$1 and password=$2", user.UserName, user.Password)

		for rows.Next() {
			rows.Scan(&user.UserName, &user.Password)
			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users)
	}
}
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		param := mux.Vars(r)
		id, _ := strconv.Atoi(param["id"])
		rows, err := db.Exec("delete from users where id=$1", id)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(rows)
	}
}
