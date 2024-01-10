package initilization

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var RedisClient *redis.Client

func ConnectToDatabase() (*gorm.DB, *redis.Client, error) {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DBNAME")
	)

	pgsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(pgsqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect successful to Database!")

	// Redis initialization
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Update with your Redis server address
		Password: "",               // Set password if applicable
		DB:       0,                // Default DB
	})

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}

	return db, RedisClient, nil
}
