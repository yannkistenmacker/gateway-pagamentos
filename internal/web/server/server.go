package server

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/go-chi/chi"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}
