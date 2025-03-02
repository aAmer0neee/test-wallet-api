package domain

import (
	"sync"

	"github.com/google/uuid"
)

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

	CachedWallets struct {	// В крупных системах лучше использовать отдельную дб для хранения кеша
		wallets sync.Map	// Например Redis
	}
)

func (w *Wallet) Deposit(amount float64) {
	if amount >= 0 {
		w.Balance += amount
	}
}

func (w *Wallet) Withdraw(amount float64) {
	if amount >= 0 && amount < w.Balance {
		w.Balance -= amount
	}
}

func InitCache()(* CachedWallets){
	return  &CachedWallets{}
}

func (cw *CachedWallets) GetWallet(id uuid.UUID) (*Wallet, bool) {
	
	value, ok := cw.wallets.Load(id)
	if !ok || value == nil {
		return nil, false
	}
	wallet, ok := value.(*Wallet)
	if !ok {
		return nil, false
	}
	return wallet, true
}

func (cw *CachedWallets) AddWallet(id uuid.UUID, wallet *Wallet) {
	if wallet == nil {
		return
	}
	cw.wallets.Store(id, wallet)
}