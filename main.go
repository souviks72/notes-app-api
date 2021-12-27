package main

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/souviks72/notes-app-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c        *mongo.Client
	db       *mongo.Database
	notesCol *mongo.Collection
	cfg      config.Properties
)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Printf("Error reading config %+v", err)
	}

	connectionURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		fmt.Printf("Error connecting to db %+v", err)
	}

	defer func() {
		err = mongoClient.Disconnect(context.TODO())
		if err != nil {
			fmt.Printf("Error disconnecting from db %+v", err)
		}
	}()

	db = mongoClient.Database(cfg.DBName)
	notesCol = db.Collection(cfg.NotesCollection)
}