package repository

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/google/uuid"
	_ "github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sql.DB
	Cache *domain.CachedWallets
}

func InitRepository(db *sql.DB, cache *domain.CachedWallets) *Repository {
	return &Repository{db: db,
	Cache: cache,}
}

func (d *Repository) GetWallet(walletID uuid.UUID) (*domain.Wallet, error) {
	
	wallet, ok := d.Cache.GetWallet(walletID)

	if !ok	{

		wallet = &domain.Wallet{}

		err := d.db.QueryRow("SELECT id, balance from wallets where id = $1",
		walletID).Scan(&wallet.ID, &wallet.Balance); if err != nil {
			return nil, fmt.Errorf("error getting balance %s",err.Error())
		}

		d.Cache.AddWallet(wallet.ID, wallet)


	}

	return wallet, nil
}

func (d *Repository) CreateWallet(ctx context.Context, walletId uuid.UUID) (*domain.Wallet, error) {

	var wallet domain.Wallet

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, fmt.Errorf("error create wallet %s", err. Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = tx.QueryRow("select id, balance from wallets where id = $1",
	 	walletId).Scan(&wallet.ID, &wallet.Balance); err == nil  {
		return &wallet, nil
	}

	if err != sql.ErrNoRows{
		return nil, fmt.Errorf("error getting wallet %s", err.Error())
	}

	if _, err = tx.Exec("insert into wallets(id) values ($1) on conflict(id) do nothing", walletId); err != nil {
		return nil, fmt.Errorf("error creating wallet %s", err.Error())
	}

	d.Cache.AddWallet(wallet.ID, &wallet)
	return &domain.Wallet{ID: walletId}, nil
}

func (d *Repository) ChangeBalance(ctx context.Context, walletId uuid.UUID, balance float64) (error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("error change wallet %s", err. Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
 	
	if _, err = tx.ExecContext(ctx, "SELECT balance FROM wallets WHERE id = $1 FOR UPDATE",walletId); err != nil {
		return fmt.Errorf("error get wallet %s", err.Error())
	}


	if _, err = tx.ExecContext(ctx, "update wallets set balance = balance + $1 where id = $2 and balance + $1 >= 0", balance, walletId); err != nil {

		return fmt.Errorf("error update wallet %s", err.Error())
	}

	return nil

}