package gapi

import (
	"fmt"

	db "github.com/blanc08/go-simple-bank/db/sqlc"
	"github.com/blanc08/go-simple-bank/pb"
	"github.com/blanc08/go-simple-bank/token"
	"github.com/blanc08/go-simple-bank/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer create a new gRPC server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
