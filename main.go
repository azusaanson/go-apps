package main

import (
	"net"

	"github.com/azusaanson/invest-api/config"
	"github.com/azusaanson/invest-api/db/db"
	"github.com/azusaanson/invest-api/gapi"
	"github.com/azusaanson/invest-api/proto/pb"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	// if config.DBDriver == "mysql"
	dbSource := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName

	conn, err := gorm.Open(mysql.Open(dbSource+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, "mysql://"+dbSource)

	store := db.NewStore(conn)

	runGrpcServer(config, store)
}

func runGrpcServer(config config.Config, store db.StoreInterface) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInvestServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
