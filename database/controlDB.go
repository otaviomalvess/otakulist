package database

import (
	"context"
	"log"
	"otavio-alves/OtakuList/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

// FindEntity .. Returns the requested entity
func FindEntity(filter interface{}, colName string) (res interface{}, err error) {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Calls DB function to find entity correspondent to the given name
	if err = Db.Collection(colName).FindOne(ctx, filter).Decode(&res); err != nil {
		log.Printf(configs.WARN_FINDING_DOC, err)
		return
	}

	return
}

// FindEntities .. Returns more than one requested entity
func FindEntities(filter interface{}, colName string, limit int64) (res []interface{}, err error) {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Creates find options
	opts := &options.FindOptions{
		Limit: &limit,
	}

	// Calls DB function to find all entities correspondent to the given filter
	cur, err := Db.Collection(colName).Find(ctx, filter, opts)
	if err != nil {
		log.Printf(configs.WARN_FINDING_DOC, err)
		return
	}

	// Decodes each document into results and
	if err = cur.All(context.TODO(), &res); err != nil {
		log.Printf(configs.WARN_DECODING_DOCS, err)
		return
	}

	return
}

// InsertEntity .. Inserts the given entity
func InsertEntity(entity interface{}, colName string) error {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Calls DB function to insert entity struct
	if _, err := Db.Collection(colName).InsertOne(ctx, entity); err != nil {
		log.Printf(configs.WARN_INSERTING_DOC, err)
		return err
	}

	return nil
}

// UpdateEntity .. Searches and updates entity
func UpdateEntity(filter interface{}, operation bson.M, colName string) error {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Calls DB function to update document
	if _, err := Db.Collection(colName).UpdateOne(ctx, filter, operation); err != nil {
		log.Printf(configs.WARN_UPDATING_DOC, err)
		return err
	}

	return nil
}

// UpdateManyEntities .. Searches and updates more than one entity
func UpdateManyEntities(filter interface{}, operation bson.M, colName string) error {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Calls DB function to update document
	if _, err := Db.Collection(colName).UpdateOne(ctx, filter, operation); err != nil {
		log.Printf(configs.WARN_UPDATING_DOC, err)
		return err
	}

	return nil
}

// DeleteEntity .. Searches and deletes entity
func DeleteEntity(filter bson.M, colName string) error {

	// Creates a context for deletion
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Calls cancel context function
	defer cancel()

	// Calls DB function do delete document
	if _, err := Db.Collection(colName).DeleteOne(ctx, filter); err != nil {
		log.Printf(configs.WARN_DELETING_DOC, err)
		return err
	}

	return nil
}
