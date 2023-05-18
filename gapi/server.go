package gapi

import (
	"github.com/azusaanson/invest-api/config"
	"github.com/azusaanson/invest-api/db/db"
	"github.com/azusaanson/invest-api/proto/pb"
)

type Server struct {
	pb.UnimplementedInvestServer
	config config.Config
	store  db.StoreInterface
}

func NewServer(config config.Config, store db.StoreInterface) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
