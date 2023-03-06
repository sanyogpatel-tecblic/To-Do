package endpoints

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	"github.com/sanyogpatel-tecblic/To-Do/controller/model"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var newId int

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")
		var task model.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		json.NewDecoder(r.Body).Decode(&task.ID)
		if err != nil {
			apierror := APIError{
				Code:    http.StatusBadRequest,
				Message: "Error parsing the body: " + err.Error(),
			}

			w.WriteHeader(apierror.Code)
			json.NewEncoder(w).Encode(apierror)
			return
		}
		if task.Tasks == "" {
			apierror := APIError{
				Code:    http.StatusBadRequest,
				Message: "Task is required",
			}
			w.WriteHeader(apierror.Code)
			json.NewEncoder(w).Encode(apierror)
			return
		}
		err = db.QueryRowContext(ctx, "INSERT INTO todo (tasks) VALUES ($1) returning id", task.Tasks).Scan(&newId)

		if err != nil {
			http.Error(w, "Error parsing request body 2", http.StatusBadRequest)
			return
		}
		if err == nil {
			task = model.Task{
				ID:         newId,
				Tasks:      task.Tasks,
				Statuscode: http.StatusCreated,
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)
		}
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
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error while fetching the task: %s", err.Error())
			return
		}

		// var task model.Task
		for rows.Next() {
			var task model.Task
			err = rows.Scan(&task.ID, &task.Tasks)
			if err != nil {
				log.Fatal(err)
				return
			}
			tasks = append(tasks, task)
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(tasks)
		}
	}
}
func GetTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		var task model.Task
		row := db.QueryRow("SELECT id,tasks FROM todo WHERE id=$1", id)

		err = row.Scan(&task.ID, &task.Tasks)
		if err != nil {
			apierror := APIError{
				Code:    http.StatusBadRequest,
				Message: "No such rows with provided id is available!!",
			}
			w.WriteHeader(apierror.Code)
			json.NewEncoder(w).Encode(apierror)
			return
		}
		if err == nil {
			task = model.Task{
				ID:         task.ID,
				Tasks:      task.Tasks,
				Statuscode: http.StatusAccepted,
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(task)
		}
	}
}
func UpdateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		TaskID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		var task model.Task

		json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			apierror := APIError{
				Code:    http.StatusBadRequest,
				Message: "Error parsing the body: " + err.Error(),
			}
			w.WriteHeader(apierror.Code)
			json.NewEncoder(w).Encode(apierror)
			return
		}
		if task.Tasks == "" {
			apierror := APIError{
				Code:    http.StatusBadRequest,
				Message: "Task is required",
			}
			w.WriteHeader(apierror.Code)
			json.NewEncoder(w).Encode(apierror)
			return
		}
		result, err := db.Exec("update todo set tasks=$1 where id=$2", task.Tasks, TaskID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error deleting task: %s", err.Error())
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error getting rows affected: %s", err.Error())
			return
		}
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Task not found with ID: %s", TaskID)
			return
		}
		if err == nil {
			// id := strconv.Itoa(TaskID)
			task = model.Task{
				ID:         TaskID,
				Tasks:      task.Tasks,
				Statuscode: http.StatusOK,
			}
			// response := map[string]string{"id": id, "message": "User updated successfully", "tasks": task.Tasks, "statuscode": http.StatusText(http.StatusOK)}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(http.StatusOK)
			json.NewEncoder(w).Encode(task)
		}
	}
}
func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		// defer cancel()

		w.Header().Set("Content-Type", "application/json")
		TaskID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		// var task model.Task
		// query := fmt.Sprintf("delete from todo where id=$1", TaskID)
		result, err := db.Exec("delete from todo where id=$1", TaskID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error deleting task: %s", err.Error())
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error getting rows affected: %s", err.Error())
			return
		}
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "task not found with ID: %s", TaskID)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task deleted successfully!")
	}
}
func MarkAsDone(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		var task model.Task

		res, err := db.Exec("update todo set done=1 where id=$1", id)
		if err != nil {
			log.Println(err)
		}
		rowsAffected, err := res.RowsAffected()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error getting rows affected: %s", err.Error())
			return
		}
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Task not found with ID: %s", id)
			return
		}
		if err == nil {
			task = model.Task{
				ID:         id,
				Tasks:      task.Tasks,
				Statuscode: http.StatusOK,
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Task Done! where id is: %s", id)
			json.NewEncoder(w).Encode("success")
		}
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
