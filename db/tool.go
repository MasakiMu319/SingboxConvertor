package db

import (
	"SingboxConvertor/api/model"
	"SingboxConvertor/utils"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var MongoClient *mongo.Client

func InitMongoClient() error {
	file, _ := os.ReadFile(DNSConfig)
	DNSs := make([]model.DNS, 0)
	err := json.Unmarshal(file, &DNSs)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(MongoURI))
	if err != nil {
		return err
	}

	if err = MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	collection := MongoClient.Database(DB).Collection(DNSCollection)

	for _, DNS := range DNSs {
		err = collection.FindOne(
			ctx, bson.M{"_id": DNS.ServerID}).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				_, err = collection.InsertOne(ctx, DNS)
				if err != nil {
					utils.ConvertorLogPrintf(err, "Insert DNS config to MongoDB failed.")
					return err
				}
			} else {
				utils.ConvertorLogPrintf(err, "Something wrong had happened.")
				return err
			}
		} else {
			utils.ConvertorLogPrintln("Skipped insert DNS config:", DNS.ServerTag)
		}
	}

	return nil
}
