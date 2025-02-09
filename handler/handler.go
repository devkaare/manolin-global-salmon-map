package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devkaare/manolin-global-salmon-map/farms"
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
