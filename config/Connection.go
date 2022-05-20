package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// package config

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func Connect() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}
// 	URL := os.Getenv("MONGO_URL")
// 	// URL := "http://localhost:27017"
// 	fmt.Println("os:", URL)

// 	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		// log.Fatal(err)

// 	}
// 	db := client.Database("goauth")
// 	fmt.Println(db)
// 	//return client
// }
