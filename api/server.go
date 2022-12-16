package api

import (
	db "ShelterChatBackend/Api/db/sqlc"
	"ShelterChatBackend/Api/token"
	"ShelterChatBackend/Api/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewSever(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	//authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))


	server.router = router
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
