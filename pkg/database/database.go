package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func MySQLConnection() *sql.DB {
	godotenv.Load()
	dsn := os.Getenv("MYSQL_DSN")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")
	return db
}

func Redisconnection() *redis.Client {

	godotenv.Load()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Redis!")
	return rdb

}
