package control

import (
	"encoding/json"
	"log"
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/database"
	"otavio-alves/OtakuList/util"
	"otavio-alves/OtakuList/validation"

	"go.mongodb.org/mongo-driver/bson"
)

// loginUser .. Lists the requested entity
func loginUser(email string, colName string) ([]byte, int, string) {

	// Calls find entity DB function using the object ID as filter and a DB collection
	entity, err := database.FindEntity(bson.M{configs.EMAIL: email}, colName)
	if err != nil {
		return nil, http.StatusUnauthorized, configs.ERR_INVALID_LOGIN
	}

	// Parses entity struct into JSON byte slice
	bytes, err := util.StructToSON(entity, json.Marshal)
	if err != nil {
		return nil, http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return bytes, 0, configs.EMPTY_STR
}

// listEntity .. Lists the requested entity
func listEntity(name string, colName string) ([]byte, int, string) {

	// Calls find entity DB function using the object ID as filter and a DB collection
	entity, err := database.FindEntity(bson.M{configs.NAME: name}, colName)
	if err != nil {
		return nil, http.StatusNotFound, configs.ERR_NOT_FOUND
	}

	// Parses entity struct into JSON byte slice
	bytes, err := util.StructToSON(entity, json.Marshal)
	if err != nil {
		return nil, http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return bytes, 0, configs.EMPTY_STR
}

// checkExistence .. Cheks if entity exists in DB
func checkExistence(name string, colName string) (bool, error) {

	// Calls find entity DB function using the object ID as filter and a DB collection
	entity, err := database.FindEntity(bson.M{configs.NAME: bson.M{"$exists": true, "$eq": name}}, colName)
	if err != nil {
		return false, err
	}
	log.Print(entity)
	return entity != nil, nil
}

// searchEntities .. Search entities by name
func searchEntities(name string, colName string, limit int64) ([]byte, int, string) {

	filter := bson.M{configs.NAME: name}
	if name == configs.EMPTY_STR {
		filter = bson.M{}
	}

	// Calls find entities DB function using the name as filter and the DB collection
	entities, err := database.FindEntities(filter, colName, limit)
	if err != nil {
		return nil, http.StatusBadRequest, configs.ERR_BAD_REQ
	}

	// Parses entities struct into JSON byte slice
	bytes, err := util.StructToSON(entities, json.Marshal)
	if err != nil {
		return nil, http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return bytes, 0, configs.EMPTY_STR
}

// filterEntity .. Search entities with a filter
func filterEntity(search interface{}, colName string) ([]byte, int, string) {

	// Calls find entities DB function using given filter and the DB collection
	res, err := database.FindEntities(search, colName, configs.LIM_FILTER)
	if err != nil {
		return nil, http.StatusBadRequest, configs.ERR_BAD_REQ
	}

	// Parses entity struct into JSON byte slice
	bytes, err := util.StructToSON(res, json.Marshal)
	if err != nil {
		return nil, http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return bytes, 0, configs.EMPTY_STR
}

// registerEntity .. Registers the given entity in the DB
func registerEntity(cols []string, insert ...interface{}) (int, string) {

	// Calls insert user DB function
	for i, colName := range cols {

		// Validates insert interface
		if err := validation.Validator.Struct(insert[i]); err != nil {
			log.Printf(configs.WARN_VALIDATION, err)
			return http.StatusBadRequest, configs.ERR_BAD_REQ
		}

		if err := database.InsertEntity(insert[i], colName); err != nil {
			return http.StatusInternalServerError, configs.ERR_INTERN_SERVER
		}
	}

	return 0, configs.EMPTY_STR
}

// updateEntity .. Updates the requested entity in the DB
func updateEntity(cols []string, filter bson.M, operation string, update ...interface{}) (int, string) {

	// Calls update entity DB function
	for i, colName := range cols {

		// Validates update interface
		err := validation.Validator.Struct(update[i])
		if err != nil {
			log.Printf(configs.WARN_VALIDATION, err)
			return http.StatusBadRequest, configs.ERR_BAD_REQ
		}

		err = database.UpdateEntity(filter, bson.M{operation: update[i]}, colName)
		if err != nil {
			return http.StatusInternalServerError, configs.ERR_INTERN_SERVER
		}
	}

	return 0, configs.EMPTY_STR
}

// updateEntityList .. Updates the requested entity list in the DB
func updateEntityList(col string, filter bson.M, operation string, update bson.M, val interface{}) (int, string) {

	// Validates update interface
	if err := validation.Validator.Struct(val); err != nil {
		log.Printf(configs.WARN_VALIDATION, err)
		return http.StatusBadRequest, configs.ERR_BAD_REQ
	}

	// Calls update entity DB function
	if err := database.UpdateEntity(filter, bson.M{operation: update}, col); err != nil {
		return http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return 0, configs.EMPTY_STR
}

// updateEntitiesList .. Updates the entities list in the DB
func updateEntitiesList(col string, filter bson.M, operation string, update bson.M, val interface{}) (int, string) {

	// Validates update interface
	if err := validation.Validator.Struct(val); err != nil {
		log.Printf(configs.WARN_VALIDATION, err)
		return http.StatusBadRequest, configs.ERR_BAD_REQ
	}

	// Calls update entity DB function
	if err := database.UpdateManyEntities(filter, bson.M{operation: update}, col); err != nil {
		return http.StatusInternalServerError, configs.ERR_INTERN_SERVER
	}

	return 0, configs.EMPTY_STR
}

// deleteEntity .. Deletes the requested entity in the DB
func deleteEntity(cols []string, filter bson.M) (int, string) {

	for _, colName := range cols {

		// Calls delete entity DB function
		if err := database.DeleteEntity(filter, colName); err != nil {
			return http.StatusInternalServerError, configs.ERR_INTERN_SERVER
		}
	}

	return 0, configs.EMPTY_STR
}
