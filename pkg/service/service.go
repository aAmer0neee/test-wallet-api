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
	wallet, err := s.repository.GetWallet(transaction.WalletId)

	if err != nil {
		wallet, err = s.repository.CreateWallet(ctx, transaction.WalletId)
		if err != nil {
			log.Println(err)
			return
		}
	}

	switch transaction.OperationType {
	case "DEPOSIT":
		if err = s.repository.ChangeBalance(ctx, wallet.ID, transaction.Amount); err != nil {
			log.Println(err)
		}
		case "WITHDRAW":
		if err = s.repository.ChangeBalance(ctx, wallet.ID, -transaction.Amount); err != nil {
			log.Println(err)
		}

	}
}
