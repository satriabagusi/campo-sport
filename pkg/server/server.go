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

// type Server struct {
// 	config     Duration
// 	tokenMaker token.Maker
// 	router     *gin.Engine
// }

// func NewServer(config Duration) (*Server, error) {
// 	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot create token maker: %w", err)
// 	}

// 	server := &Server{
// 		config:     config,
// 		tokenMaker: tokenMaker,
// 	}

// 	return server, nil
// }

// func (s *Server) Initialize(connstr string) error {
// 	// initialize DB
// 	db, err := sql.Open("postgres", connstr)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	//initialize repo
// 	userRepo := repository.NewUserRepository(db)
// 	courtRepo := repository.NewCourtRepository(db)
// 	bookingRepo := repository.NewBookingRepository(db)
// 	voucherRepo := repository.NewVoucherRepository(db)
// 	userDetailRepo := repository.NewUserDetailRepository(db)

// 	//initialize usecase
// 	userUsecase := usecase.NewUserUsecase(userRepo, userDetailRepo)
// 	courtUsecase := usecase.NewCourtUsecase(courtRepo)
// 	bookingUsecase := usecase.NewBookingUsecase(bookingRepo)
// 	voucherUsecase := usecase.NewVoucherUsecase(voucherRepo)
// 	userDetailUsecase := usecase.NewUserDetailUsecase(userDetailRepo)

// 	//setup router
// 	r := gin.Default()
// 	api := r.Group("/api/v1")
// 	router.NewUserRouter(api, userUsecase)
// 	router.NewCourtRouter(api, courtUsecase)
// 	router.NewBookingRouter(api, bookingUsecase)
// 	router.NewVoucherRouter(api, voucherUsecase)
// 	router.NewUserDetailRouter(api, userDetailUsecase, userUsecase)

// 	s.router = r
// 	return nil

// }

// func (s *Server) Start(addr string) error {
// 	return http.ListenAndServe(addr, s.router)
// }

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

	//setup router
	r := gin.Default()
	api := r.Group("/api/v1")
	router.NewUserRouter(api, userUsecase)
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
