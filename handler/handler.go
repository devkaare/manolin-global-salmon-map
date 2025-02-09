package handler

import (
	// "context"
	// "encoding/json"
	// "errors"
	// "log"
	// "math/rand/v2"
	"encoding/json"
	"log"
	"net/http"

	"github.com/devkaare/manolin-global-salmon-map/farms"
	// "strconv"
	// "github.com/go-chi/chi/v5"
)

type New struct {
}

func (t *New) Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (t *New) Farms(w http.ResponseWriter, r *http.Request) {
	farms, err := farms.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, _ := json.Marshal(farms)
	_, _ = w.Write(jsonResp)
}

// func (t *Todo) Health(w http.ResponseWriter, r *http.Request) {
// 	jsonResp, _ := json.Marshal(t.Repo.Health())
// 	_, _ = w.Write(jsonResp)
// }

// func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
// 	todo := &model.Todo{
// 		ID:          rand.Uint32N(2147483647),
// 		Title:       r.FormValue("title"),
// 		Description: r.FormValue("description"),
// 	}
// 	if _, err := t.Repo.GetTodoByID(todo.ID); err == errors.New("todo not found") && err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	if err := t.Repo.CreateTodo(todo); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	views.TodoPost(todo).Render(context.Background(), w)
// }
