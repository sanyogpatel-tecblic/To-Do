package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanyogpatel-tecblic/To-Do/controller/endpoints"
)

func Router() {
	db, err := sql.Open("postgres", "postgresql://postgres:root@localhost/todo?sslmode=disable")
	fmt.Println("connected to datbase!")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", endpoints.GetAllTasks(db)).Methods("GET")
	router.HandleFunc("/tasks/{id}", endpoints.GetTask(db)).Methods("GET")
	router.HandleFunc("/donetasks", endpoints.GetDoneTasks(db)).Methods("GET")
	router.HandleFunc("/tasks", endpoints.CreateTask(db)).Methods("POST")
	router.HandleFunc("/tasks/{id}", endpoints.UpdateTask(db)).Methods("PUT")
	router.HandleFunc("/tasks/{id}", endpoints.DeleteTask(db)).Methods("DELETE")
	router.HandleFunc("/tasks/done/{id}", endpoints.MarkAsDone(db)).Methods("PUT")

	router.HandleFunc("/users", endpoints.GetAllUsers(db)).Methods("GET")
	router.HandleFunc("/addusers", endpoints.RegisterUsers(db)).Methods("POST")
	router.HandleFunc("/updateusers/{id}", endpoints.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/login", endpoints.Login(db)).Methods("POST")
	router.HandleFunc("/deleteusers/{id}", endpoints.DeleteUser(db)).Methods("DELETE")
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8010 ...")
	http.ListenAndServe(":8010", router)
}
