package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// MangaList .. Gets user manga list in DB
func MangaList(w http.ResponseWriter, r *http.Request) {

	// Gets user by ID in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_MANGA_LIST)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// AddMangaInList .. Adds manga in user manga list
func AddMangaInList(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	var manga model.MangaInList

	if err := util.BodyToModel(r.Body, &manga); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	filter := bson.M{configs.GEN_ID: c.Value}
	add := bson.M{configs.LIST: manga}

	// Inserts manga in list
	if statusCode, err := updateEntityList(configs.COL_MANGA_LIST, filter, configs.OP_PUSH, add, manga); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// UpdateMangaInList .. Updates manga valuation in user manga list
func UpdateMangaInList(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	var val model.ContentValuation

	if err := util.BodyToModel(r.Body, &val); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	filter := bson.M{configs.GEN_ID: c.Value, configs.OP_LIST_NAME: path.Base(r.URL.Path)}
	upd := bson.M{configs.OP_LIST_VALUATION_UPDATE: val}

	// Updates manga valuation in list
	if statusCode, err := updateEntityList(configs.COL_MANGA_LIST, filter, configs.OP_SET, upd, val); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_UPDATING_ENTITY))
}

// DeleteMangaInList .. Deletes manga from user manga list
func DeleteMangaInList(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var manga model.MangaInList

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &manga); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	// Deletes manga from list
	if statusCode, err := updateEntityList(configs.COL_MANGA_LIST, bson.M{configs.GEN_ID: c.Value},
		configs.OP_PULL, bson.M{configs.LIST: manga}, manga); statusCode != 0 {

		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
