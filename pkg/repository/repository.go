package repository

import (
	"database/sql"
	"fmt"

	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/google/uuid"
	_ "github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB )(*Repository){
	return &Repository{db: db}
}

func (d *Repository) GetWallet(walletID uuid.UUID) (*domain.Wallet, error) {
	var wallet domain.Wallet
	response := d.db.QueryRow("SELECT id, balance from wallets where id = $1", walletID)

	err := response.Scan(&wallet.ID, &wallet.Balance)

	if err != nil {
		return &domain.Wallet{}, err
	}

	return &wallet, nil
}

func (d *Repository) ChangeBalance(walletId uuid.UUID, balanceValue float64) {
	_, err := d.db.Exec("update wallets set balance = $1 where id = $2", balanceValue, walletId)

	if err != nil {
		fmt.Println(err)
	}
}
