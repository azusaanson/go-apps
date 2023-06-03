package gapi

import (
	"context"
	"fmt"

	"github.com/azusaanson/invest-api/config"
	"github.com/azusaanson/invest-api/db/db"
	"github.com/azusaanson/invest-api/domain"
	"github.com/azusaanson/invest-api/proto/pb"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Server struct {
	pb.UnimplementedInvestServer
	config     config.Config
	store      db.StoreInterface
	tokenMaker domain.TokenMaker
}

func NewServer(config config.Config, store db.StoreInterface) (*Server, error) {
	symmetricKey, err := domain.NewSymmetricKeyFromString(config.TokenSymmetricKey)
	if err != nil {
		return nil, serverError(err)
	}

	tokenMaker, err := domain.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, serverError(fmt.Errorf("cannot create token maker: %w", err))
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

func (server *Server) extractMetadata(ctx context.Context) (*domain.UserMetaData, error) {
	var userAgent domain.UserAgent
	var clientIp domain.ClientIp
	var err error = nil

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			userAgent, err = domain.NewUserAgent(userAgents[0])
		}

		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			userAgent, err = domain.NewUserAgent(userAgents[0])
		}

		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			clientIp, err = domain.NewClientIp(clientIPs[0])
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		clientIp, err = domain.NewClientIp(p.Addr.String())
	}

	userMetadata, err := domain.NewUserMetadata(userAgent, clientIp)
	if err != nil {
		return nil, serverError(err)
	}

	return userMetadata, nil
}
