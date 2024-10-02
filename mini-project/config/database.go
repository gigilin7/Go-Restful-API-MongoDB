package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	// 載入 .env 檔案
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 從環境變數中取得密碼
	password := os.Getenv("MONGODB_PASSWORD")
	if password == "" {
		log.Fatal("MONGODB_PASSWORD not set in .env file")
	}

	// 建立 MongoDB 連線字串
	connPattern := fmt.Sprintf("mongodb+srv://7111056013:%s@cluster0.hztxv.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", password)

	// 設定 MongoDB 連線選項
	clientOptions := options.Client().ApplyURI(connPattern)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// 建立 Context 並設定連線 Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 連接 MongoDB
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// 回傳指定的 Database
	return client.Database("mydb"), nil
}
