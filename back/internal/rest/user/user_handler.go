package user

import (
	"encoding/json"
	"net/http"
	"tbd/internal/database/models"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)
type UserHandler struct {
	DB *bun.DB
}

func (b UserHandler) ListUser(w http.ResponseWriter, r *http.Request)  {

}
func (b UserHandler) GetUser(w http.ResponseWriter, r *http.Request)    {
	ctx := r.Context()
	id := chi.URLParam(r, "id")	

	user := new(models.User)
	err := b.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if (err != nil) {
		panic(nil)
	}

	err = json.NewEncoder(w).Encode(user)
	if (err != nil) {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

func (b UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := new(models.User)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = b.DB.NewInsert().Model(user).Exec(ctx)
	if (err != nil) {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (b UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}