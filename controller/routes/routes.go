package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sanyogpatel-tecblic/To-Do/controller/config"
	"github.com/sanyogpatel-tecblic/To-Do/controller/endpoints"
)

func Routes(app *config.AppConfig) http.Handler {
	db, err := sql.Open("postgres", "postgresql://postgres:root@localhost/todo?sslmode=disable")
	fmt.Println("connected to datbase!")
	if err != nil {
		log.Fatal(err)
	}
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/tasks", endpoints.GetAllTasks(db))
	mux.Get("/tasks/{id}", endpoints.GetTask(db))
	mux.Get("/donetasks", endpoints.GetDoneTasks(db))
	mux.Post("/create/tasks", endpoints.CreateTask(db))
	mux.Put("/tasks/{id}", endpoints.UpdateTask(db))
	mux.Delete("/deletetasks/{id}", endpoints.DeleteTask(db))
	mux.Put("/tasks/done/{id}", endpoints.MarkAsDone(db))

	//----------	mux.Get("/tasks", endpoints.GetAllTasks(db))
	mux.Get("/users", endpoints.GetAllUsers(db))
	mux.Post("/users", endpoints.RegisterUsers(db))
	mux.Post("/updateusers/{id}", endpoints.UpdateUser(db))
	mux.Put("/users/login", endpoints.Login(db))
	mux.Delete("/deleteusers/{id}", endpoints.DeleteUser(db))
	mux.Put("/tasks/done/{id}", endpoints.MarkAsDone(db))

	return mux
}

// func Router() {
// db, err := sql.Open("postgres", "postgresql://postgres:root@localhost/todo?sslmode=disable")
// fmt.Println("connected to datbase!")
// if err != nil {
// 	log.Fatal(err)
// }
// 	defer db.Close()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/tasks", endpoints.GetAllTasks(db)).Methods("GET")
// 	router.HandleFunc("/tasks/{id}", endpoints.GetTask(db)).Methods("GET")
// 	router.HandleFunc("/donetasks", endpoints.GetDoneTasks(db)).Methods("GET")
// 	router.HandleFunc("/tasks", endpoints.CreateTask(db)).Methods("POST")
// 	router.HandleFunc("/tasks/{id}", endpoints.UpdateTask(db)).Methods("PUT")
// 	router.HandleFunc("/tasks/{id}", endpoints.DeleteTask(db)).Methods("DELETE")
// 	router.HandleFunc("/tasks/done/{id}", endpoints.MarkAsDone(db)).Methods("PUT")

// 	router.HandleFunc("/users", endpoints.GetAllUsers(db)).Methods("GET")
// 	router.HandleFunc("/addusers", endpoints.RegisterUsers(db)).Methods("POST")
// 	router.HandleFunc("/updateusers/{id}", endpoints.UpdateUser(db)).Methods("PUT")
// 	router.HandleFunc("/users/login", endpoints.Login(db)).Methods("POST")
// 	router.HandleFunc("/deleteusers/{id}", endpoints.DeleteUser(db)).Methods("DELETE")

// }
