package gapi

import (
	"github.com/azusaanson/invest-api/config"
	"github.com/azusaanson/invest-api/db/store"
	"github.com/azusaanson/invest-api/proto/pb"
)

type Server struct {
	pb.UnimplementedInvestServer
	config config.Config
	store  *store.Store
}

func NewServer(config config.Config, store *store.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
