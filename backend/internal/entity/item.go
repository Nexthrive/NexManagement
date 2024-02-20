package entity

import "time"

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Deskripsi string    `json:"Deskripsi"`
	CreatedAt time.Time `json:"created_at"`
}
