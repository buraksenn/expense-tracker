// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"time"
)

type Expense struct {
	ID        int32     `db:"id" json:"id"`
	UserID    int32     `db:"user_id" json:"userID"`
	Name      string    `db:"name" json:"name"`
	Type      string    `db:"type" json:"type"`
	Price     float32   `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type User struct {
	ID        int32     `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	ChatID    string    `db:"chat_id" json:"chatID"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}