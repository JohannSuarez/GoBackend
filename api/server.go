package api

import (
    "github.com/gin-gonic/gin"
    db "github.com/JohannSuarez/GoBackend/db/sqlc"
)

// SQLStore provides all functions to execute SQL queries and transactions
type Server struct {
    store db.Store
    router *gin.Engine
}


// NewServer creates a new HTTP server and setup routing
func NewStore(store db.Store) *Server {

    server := &Server{store: store}
    router := gin.Default()

    // Add routes to router
    router.POST("/accounts", server.createAccount)
    router.GET("/accounts/:id", server.getAccount)
    router.GET("/accounts", server.listAccount)

    server.router = router
    return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}
