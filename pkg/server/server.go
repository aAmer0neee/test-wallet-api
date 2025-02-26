package server

import (
	"encoding/json"
	"fmt"

	"github.com/aAmer0neee/test-wallet-api/pkg/domain"
	"github.com/aAmer0neee/test-wallet-api/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	Server struct {
		router  *gin.Engine
		service *service.Service
	}
)

func InitServer(service *service.Service)(*Server){
	return &Server{
		router:  gin.Default(),
		service: service,
	}
}

func (s *Server) Up(port string) {

	s.initRoutes()

	gin.SetMode(gin.ReleaseMode)
	s.router.Run(port)
}

func (s *Server) initRoutes() {

	s.router.GET("/api/v1/wallets/:WALLET_UUID", s.inquiryBalance)

	s.router.POST("api/v1/wallet", s.changeBalance)
}

func (s *Server) inquiryBalance(ctx *gin.Context) {

	walletID, _ := uuid.Parse(ctx.Param("WALLET_UUID"))
	rBody := domain.Transaction{
		WalletId:      walletID,
		OperationType: "INQUIRY",
		Amount:        0,
	}

	wallet := s.service.InquiryWallet(rBody)
	balance := json.NewEncoder(ctx.Writer)
	balance.SetIndent("", "  ")
	balance.Encode(wallet)
}


func (s *Server) changeBalance(ctx *gin.Context) {
	requestBody := domain.Transaction{}

	ctx.ShouldBindBodyWithJSON(&requestBody)

	switch requestBody.OperationType {

	case "DEPOSIT":
		s.service.DepositWallet()

	case "WITHDRAW":
		s.service.WithdrawWallet()

	default:
		fmt.Println("invalid operation tipe for change wallet method")

	}
}
