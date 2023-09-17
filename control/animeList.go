package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// AnimeList .. Gets user anime list in DB
func AnimeList(w http.ResponseWriter, r *http.Request) {

	// Gets user by ID in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_ANIME_LIST)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// AddAnimeInList .. Adds anime in user anime list
func AddAnimeInList(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var anime model.AnimeInList

	if err := util.BodyToModel(r.Body, &anime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	filter := bson.M{configs.GEN_ID: c.Value}
	add := bson.M{configs.LIST: anime}

	// Inserts anime in list
	if statusCode, err := updateEntityList(configs.COL_ANIME_LIST, filter, configs.OP_PUSH, add, anime); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// UpdateAnimeInList .. Updates anime valuation in user anime list
func UpdateAnimeInList(w http.ResponseWriter, r *http.Request) {

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

	// Updates anime valuation in list
	if statusCode, err := updateEntityList(configs.COL_ANIME_LIST, filter, configs.OP_SET, upd, val); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_UPDATING_ENTITY))
}

// DeleteAnimeInList .. Deletes anime from user anime list
func DeleteAnimeInList(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var anime model.AnimeInList

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &anime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	// Deletes anime from list
	if statusCode, err := updateEntityList(configs.COL_ANIME_LIST, bson.M{configs.GEN_ID: c.Value},
		configs.OP_PULL, bson.M{configs.LIST: anime}, anime); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
