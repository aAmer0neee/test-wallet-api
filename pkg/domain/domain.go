package domain

import "github.com/google/uuid"

type (
	Transaction struct {
		WalletId      uuid.UUID `json:"walletId"`
		OperationType string    `json:"operationType"`
		Amount        float64   `json:"amount"`
	}

	Wallet struct {
		ID      uuid.UUID `json:"walletId"`
		Balance float64   `json:"balance"`
	}
)

func (w *Wallet) Deposit(amount float64) {
	if amount > 0 {
		w.Balance += amount
	}
}

func (w *Wallet) Withdraw(amount float64) {
	if amount > 0 && amount < w.Balance {
		w.Balance -= amount
	}
}
