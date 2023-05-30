package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validator := validator.New()

	//initialize repo
	userRepo := repository.NewUserRepository(db)
	userTopUpRepo := repository.NewUserTopUpRepository(db)

	courtRepo := repository.NewCourtRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	voucherRepo := repository.NewVoucherRepository(db)
	userDetailRepo := repository.NewUserDetailRepository(db)

	//initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo, userDetailRepo, validator)
	courtUsecase := usecase.NewCourtUsecase(courtRepo, validator)
	bookingUsecase := usecase.NewBookingUsecase(bookingRepo)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo, validator)
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
	router.NewUserDetailRouter(api, userDetailUsecase, userUsecase)

	s.router = r
	return nil

}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
