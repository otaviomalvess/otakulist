package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// ListForum .. Lists the requested forum
func ListForum(w http.ResponseWriter, r *http.Request) {

	// Gets forum in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_FORUMS)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// RegisterForum .. Registers the given forum in the DB
func RegisterForum(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var forum model.Forum

	if err := util.BodyToModel(r.Body, &forum); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Inserts forum in DB
	if statusCode, err := registerEntity([]string{configs.COL_FORUMS}, forum); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// DeleteForum .. Deletes forum
func DeleteForum(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Deletes forum from DB
	if statusCode, err := deleteEntity([]string{configs.COL_FORUMS}, bson.M{configs.NAME: path.Base(r.URL.Path)}); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
