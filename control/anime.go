package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// ListAnime .. Lists the requested anime
func ListAnime(w http.ResponseWriter, r *http.Request) {

	// Gets anime by ID in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_ANIMES)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// SearchAnimes .. Search animes by name
func SearchAnimes(w http.ResponseWriter, r *http.Request) {

	// Searches for anime by name in DB
	bytes, statusCode, err := searchEntities(r.URL.Query().Get(configs.ANIME), configs.COL_SEARCH_ANIMES, configs.LIM_SEARCH_ANIME)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// FilterAnimes .. Search animes with a filter
func FilterAnimes(w http.ResponseWriter, r *http.Request) {

	b := util.MapToBSON(r.URL.Query())

	// Searches for animes by filter in DB
	bytes, statusCode, err := filterEntity(b, configs.COL_FILTER_ANIMES)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// RegisterAnime .. Registers the given anime in the DB
func RegisterAnime(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var anime model.RegistrableAnime

	if err := util.BodyToModel(r.Body, &anime.Anime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	// Checks if anime already exists in DB
	if b, _ := checkExistence(anime.Anime.Name, configs.COL_ANIMES); b {
		w.WriteHeader(http.StatusUnauthorized)
		util.WriteResponse(w, []byte(configs.ERR_EXIST_ENTITY))
		return
	}

	anime.PopulateRegistrableAnime()

	// Inserts anime in DB
	if statusCode, err := registerEntity(configs.ANIMES_COLS, anime.Anime, anime.Search, anime.Filter); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// UpdateAnime .. Updates anime info
func UpdateAnime(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	var upd model.UpdatableAnime

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &upd.Anime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	upd.PopulateUpdatableAnime()
	filter := bson.M{configs.NAME: path.Base(r.URL.Path)}

	// Updates anime in DB
	if statusCode, err := updateEntity(configs.ANIMES_COLS, filter, configs.OP_SET, upd.Anime, upd.Search, upd.Filter); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	filter = bson.M{configs.OP_LIST_NAME: path.Base(r.URL.Path)}
	listUpd := bson.M{configs.OP_LIST_NAME_UPDATE: upd.InList.Name, configs.OP_LIST_TYPE_UPDATE: upd.InList.Type}

	// Updates anime in DB
	if statusCode, err := updateEntitiesList(configs.COL_ANIME_LIST, filter, configs.OP_SET, listUpd, upd.InList); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_UPDATING_ENTITY))
}

// DeleteAnime .. Deletes anime from DB
func DeleteAnime(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	// Deletes anime from DB
	if statusCode, err := deleteEntity(configs.ANIMES_COLS, bson.M{configs.NAME: path.Base(r.URL.Path)}); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
