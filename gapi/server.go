package gapi

import (
	"fmt"

	db "github.com/best-bank-api/db/sqlc"
	"github.com/best-bank-api/pb"
	"github.com/best-bank-api/token"
	"github.com/best-bank-api/util"
)

type Server struct {
	pb.UnimplementedBestBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

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
