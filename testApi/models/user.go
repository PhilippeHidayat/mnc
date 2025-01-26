package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID `pg:"user_id,pk"`
	FirstName   string    `pg:"first_name"`
	LastName    string    `pg:"last_name"`
	PhoneNumber string    `pg:"phone_number"`
	Address     string    `pg:"address"`
	Pin         string    `pg:"pin"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at,default:now()"`
}

func (u *User) BeforeInsert() error {
	u.UserID = uuid.New()
	return nil
}

func (u *User) BeforeUpdate() error {
	u.UpdatedAt = time.Now()
	return nil
}
