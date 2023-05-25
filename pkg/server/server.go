package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/satriabagusi/campo-sport/internal/delivery/router"
	"github.com/satriabagusi/campo-sport/internal/repository"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize(connstr string) error {
	// initialize DB
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	//initialize repo
	userRepo := repository.NewUserRepository(db)
	userTopUpRepo := repository.NewUserTopUpRepository(db)

	courtRepo := repository.NewCourtRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	voucherRepo := repository.NewVoucherRepository(db)
	userDetailRepo := repository.NewUserDetailRepository(db)

	//initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo, userDetailRepo)
	courtUsecase := usecase.NewCourtUsecase(courtRepo)
	bookingUsecase := usecase.NewBookingUsecase(bookingRepo)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo)
	userDetailUsecase := usecase.NewUserDetailUsecase(userDetailRepo)

	userTopUpUsecase := usecase.NewUserTopUpUsecase(userTopUpRepo)

	//setup router
	r := gin.Default()
	api := r.Group("/api/v1")
	router.NewUserRouter(api, userUsecase)
	router.NewUserTopUpRouter(api, userTopUpUsecase)
	router.NewCourtRouter(api, courtUsecase)
	router.NewBookingRouter(api, bookingUsecase)
	router.NewVoucherRouter(api, voucherUsecase)
	router.NewUserDetailRouter(api, userDetailUsecase)

	s.router = r
	return nil

}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
