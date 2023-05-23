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

	//initialize repo and usecase....
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	//setup router
	r := gin.Default()
	api := r.Group("/api/v1")
	router.NewUserRouter(api, userUsecase)

	s.router = r
	return nil

}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
