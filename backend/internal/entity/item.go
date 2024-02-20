package entity

import "time"

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Deskripsi string    `json:"Deskripsi"`
	Qty       int       `json:"qty"`
	Harga     float64   `json:"harga"`
	CreatedAt time.Time `json:"created_at"`
}
