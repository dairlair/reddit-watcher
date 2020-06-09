package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw/reddit"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	uri string
	client *mongo.Client
}

func newMongoRepository (uri string) *mongoRepository {
	return &mongoRepository{
		uri: uri,
		client: nil,
	}
}

func (mr *mongoRepository) connect () error {
	clientOptions := options.Client().ApplyURI(mr.uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(err)
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	mr.client = client
	return nil
}

func (mr *mongoRepository) save(post reddit.Post) error {
	log.Info("Post saved successfully")
	return nil
}