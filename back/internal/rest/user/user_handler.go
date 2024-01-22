package user

import "net/http"

type UserHandler struct {
	// id of some sort
	// name of some sort
}

func (b UserHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b UserHandler) GetBook(w http.ResponseWriter, r *http.Request)    {}
func (b UserHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b UserHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b UserHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}