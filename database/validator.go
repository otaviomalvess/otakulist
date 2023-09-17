package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FindForValidation .. Returns true if the given name exists in the given collection
func FindForValidation(name string, colName string) bool {

	// Creates a context for search
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	// Calls database function to find entity correspondent to the given name
	res, err := Db.Collection(colName).FindOne(ctx, bson.M{"name": name}).DecodeBytes()

	// Checks if any error occur durings
	if err != nil {
		return false
	}

	// Returns true if the result is equal to the name
	return res != nil
}
