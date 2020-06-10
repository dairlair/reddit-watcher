package storage

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw/reddit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type config struct {
	// MongoDB uri
	uri string
	// MongoDB database
	database string
}

type StorageComponent struct {
	config config
	client *mongo.Client
}

func NewStorageComponent () StorageComponent {
	config := readConfig()
	return StorageComponent{config: config}
}

// Run function implements the `github.com/stepsisters/kgb.ComponentInterface`
func (s *StorageComponent) Run () (stop func(), wait func() error, err error) {
	ch := make(chan bool, 1)
	return func() {
			ch <- true
		}, func() error {
			<- ch
			s.disconnect()
			return nil
	}, nil
}


func (s *StorageComponent) connect () error {
	clientOptions := options.Client().ApplyURI(s.config.uri)
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

	s.client = client
	return nil
}

func (s *StorageComponent) SavePost(post reddit.Post) error {
	client, err := s.getClient()
	if err != nil {
		return err
	}

	result, err := client.Database(s.config.database).Collection("posts").InsertOne(context.TODO(), post)
	if err != nil {
		log.Errorf("Post inserting error: %s", err)
	}

	log.Infof("Post saved successfully with ID: %s", result.InsertedID)
	return nil
}

func (s *StorageComponent) getClient() (*mongo.Client, error) {
	if s.client == nil {
		if err := s.connect(); err != nil {
			return nil, err
		}
	}

	return s.client, nil
}
func (s *StorageComponent) disconnect () {
	if s.client != nil {
		if err := s.client.Disconnect(context.TODO()); err != nil {
			log.Error(err)
		} else {
			log.Info("MongoDB connection closed successfully")
		}
	}
}

func readConfig() config {
	return config{
		uri:      os.Getenv("MONGODB_URI"),
		database: os.Getenv("MONGODB_DATABASE"),
	}
}