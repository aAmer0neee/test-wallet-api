package service

import (
	"context"
	"log"

	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/aAmer0neee/test-wallet-api/pkg/repository"
)

type Service struct {
	repository *repository.Repository
}

func InitService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) InquiryWallet(ctx context.Context, transaction domain.Transaction) *domain.Wallet {

	wallet, err := s.repository.GetWallet( transaction.WalletId)

	if err != nil {

		log.Println(err)
		return nil
	}

	return wallet
}

func (s *Service) ChangeWallet(ctx context.Context, transaction domain.Transaction) {
	var err error
	wallet, ok := s.repository.Cache.GetWallet(transaction.WalletId)

	if !ok {
		wallet, err = s.repository.CreateWallet(ctx, transaction.WalletId)
		if err != nil {
			log.Println(err)
			return
		}
		if wallet != nil{
		s.repository.Cache.AddWallet(wallet.ID, wallet)
		}
	}

	

	switch transaction.OperationType {
	case "DEPOSIT":
		if err = s.repository.ChangeBalance(ctx, wallet.ID, transaction.Amount); err != nil {
			log.Println(err)
		}
		/* wallet.Deposit(transaction.Amount) */
		case "WITHDRAW":
		if err = s.repository.ChangeBalance(ctx, wallet.ID, -transaction.Amount); err != nil {
			log.Println(err)
		}

		/* wallet.Withdraw(transaction.Amount) */
	}

	/* if err = s.repository.ChangeBalance(ctx, wallet.ID, wallet.Balance); err != nil {
		log.Println(err)
	} */
}
