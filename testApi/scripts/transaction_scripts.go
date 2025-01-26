package scripts

import (
	"mnc/testApi/models"

	"github.com/go-pg/pg/v10"
)

// InsertTransaction inserts a new transaction into the database
func InsertTransaction(db *pg.DB, transaction *models.Transaction) error {
	_, err := db.Model(transaction).Insert()
	return err
}
