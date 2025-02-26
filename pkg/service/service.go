package service

import (
	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	_ "github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/aAmer0neee/test-wallet-api/pkg/repository"

	/* "github.com/gin-gonic/gin" */
)

type Service struct {
	repository *repository.Repository
}

func InitService(repo *repository.Repository)(*Service){
	return &Service{repository: repo}
}

func (s *Service) InquiryWallet(transaction domain.Transaction)(*domain.Wallet) {

	wallet, _ := s.repository.GetWallet(transaction.WalletId)

	return wallet
}

func (s *Service) DepositWallet() {
	
}

func (s *Service) WithdrawWallet() {

}
