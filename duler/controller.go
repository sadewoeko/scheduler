package duler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type DulerController struct {
	Service UseCase
}

func (c *DulerController) Find(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	idDuler, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := c.Service.Find(idDuler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (c *DulerController) FindAll(w http.ResponseWriter, r *http.Request) {

	data, err := c.Service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (c *DulerController) Store(w http.ResponseWriter, r *http.Request) {

	var duler Duler

	if err := json.NewDecoder(r.Body).Decode(&duler); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := c.Service.Store(duler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
