package cockroach

import (
	"context"
	"fmt"
	config "gocourseCRUD/config/environments"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var connectionPool *pgxpool.Pool

func GetConnectionPool(ctx context.Context) *pgxpool.Pool {
	if connectionPool != nil {
		return connectionPool
	}
	log.Println("Connecting to database...")

	pool, err := pgxpool.New(ctx, GetDatabaseUrl())
	if err != nil {
		log.Panic("Couldn't create connection pool", err)
	}
	connectionPool = pool
	
	return connectionPool
}

func GetDatabaseUrl() string {
	fmt.Println(config.NewConfig().GetConfig())
	dbUrl := config.NewConfig().GetConfig().DatabaseUrl

	if dbUrl == "" {
		log.Panic("No database url specified")
	}
	return dbUrl
}
