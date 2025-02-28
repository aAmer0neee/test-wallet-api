package server

import (
	"log"
	"net/http"

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

func InitServer(service *service.Service) *Server {
	return &Server{
		router:  gin.Default(),
		service: service,
	}
}

func (s *Server) Up(port string) {

	s.initRoutes()

	gin.SetMode(gin.ReleaseMode)

	if err := s.router.Run(port); err != nil {
		log.Fatalf("error %v", err)
	}

}

func (s *Server) initRoutes() {

	s.router.GET("/api/v1/wallets/:WALLET_UUID", s.inquiryHandler)

	s.router.POST("api/v1/wallet", s.changeHandler)
}

func (s *Server) inquiryHandler(ctx *gin.Context) {

	walletID, _ := uuid.Parse(ctx.Param("WALLET_UUID"))
	requestBody := domain.Transaction{
		WalletId:      walletID,
		OperationType: "INQUIRY",
		Amount:        0,
	}

	wallet := s.service.InquiryWallet(ctx.Request.Context(), requestBody)

	ctx.JSON(http.StatusOK, wallet)
}

func (s *Server) changeHandler(ctx *gin.Context) {
	requestBody := domain.Transaction{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("errorr binding request %s", err.Error())
	}

	switch requestBody.OperationType {

	case "DEPOSIT", "WITHDRAW":
		s.service.ChangeWallet(ctx.Request.Context(), requestBody)
	default:
		log.Printf("invalid change wallet method '%s'", requestBody.OperationType)
	}
}
