package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tbd/internal/database/models"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)
type UserHandler struct {
	DB *bun.DB
}

func (b UserHandler) ListUser(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	var users []models.User

	// .Exec(ctx) returns null for some odd reason
	err := b.DB.NewSelect().Model(&users).Scan(ctx)
	if err != nil { panic(err) }

	err = json.NewEncoder(w).Encode(users)
	if err != nil {	http.Error(w, err.Error(), http.StatusBadRequest); return	}
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

func (b UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := new(models.User)
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil { http.Error(w, err.Error(), http.StatusBadRequest); return }

	if  user.ID == 0 { 
		user.ID, err = strconv.ParseInt(id, 10, 64) 
		if err != nil { panic(err) }
	}

	_, err = b.DB.NewUpdate().Model(user).WherePK().Exec(ctx)
	if err != nil { panic(err) }

	err = json.NewEncoder(w).Encode(user)
	if err != nil { http.Error(w, err.Error(), http.StatusBadRequest); return }
}
func (b UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: add a http reponse if id is not appropriate
	ctx := r.Context()
	id := chi.URLParam(r, "id")	
	
	_, err := b.DB.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil { panic(err) }	else { w.Write([]byte("User deleted")) }
}
