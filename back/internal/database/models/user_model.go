package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID int64 `json:"id" bun:",pk,autoincrement"`
	Name string `json:"name"`
}