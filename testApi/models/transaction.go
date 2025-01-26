package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   uuid.UUID `pg:"transaction_id,pk"`
	UserID          uuid.UUID `pg:"user_id"`
	Balance         float64   `pg:"balance"`
	Status          string    `pg:"status"`
	TransactionType string    `pg:"transaction_type"`
	Amount          float64   `pg:"amount"`
	Remarks         string    `pg:"remarks"`
	BalanceBefore   float64   `pg:"balance_before"`
	BalanceAfter    float64   `pg:"balance_after"`
	CreatedDate     time.Time `pg:"created_date,default:now()"`
}

func (t *Transaction) BeforeInsert(db *pg.DB) error {
	t.TransactionID = uuid.New()
	return nil
}
