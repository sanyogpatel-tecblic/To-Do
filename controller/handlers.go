package controller

// var Repo *Repository
// var DB *sql.DB

// // Repository is the repository type
// type Repository struct {
// 	App *config.AppConfig
// }

// func NewRepo(a *config.AppConfig) *Repository {
// 	return &Repository{
// 		App: a,
// 	}
// }
// func NewHandlers(r *Repository) {
// 	Repo = r
// }

// func (m *Repository) GetAllTasks(w http.ResponseWriter, r *http.Request) {
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
