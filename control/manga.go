package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// ListManga .. Lists the requested manga
func ListManga(w http.ResponseWriter, r *http.Request) {

	// Gets manga by ID in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_MANGAS)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// SearchMangas .. Searches mangas by name
func SearchMangas(w http.ResponseWriter, r *http.Request) {

	// Searches for mangas by name in DB
	bytes, statusCode, err := searchEntities(r.URL.Query().Get(configs.NAME), configs.COL_SEARCH_MANGAS, configs.LIM_SEARCH_MANGA)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// FilterMangas .. Searches mangas with a filter
func FilterMangas(w http.ResponseWriter, r *http.Request) {

	b := util.MapToBSON(r.URL.Query())

	// Searches for mangas by filter in DB
	bytes, statusCode, err := filterEntity(b, configs.COL_FILTER_MANGAS)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// RegisterManga .. Registers the given manga in the DB
func RegisterManga(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	var manga model.RegistrableManga

	if err := util.BodyToModel(r.Body, &manga.Manga); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	// Checks if anime already exists in DB
	if b, _ := checkExistence(manga.Manga.Name, configs.COL_MANGAS); b {
		w.WriteHeader(http.StatusUnauthorized)
		util.WriteResponse(w, []byte(configs.ERR_EXIST_ENTITY))
		return
	}

	manga.PopulateRegistrableManga()

	// Inserts manga in DB
	if statusCode, err := registerEntity(configs.MANGAS_COLS, manga.Manga, manga.Search, manga.Filter); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// UpdateManga .. Updates manga info
func UpdateManga(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	var upd model.UpdatableManga

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &upd.Manga); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	upd.PopulateUpdatableManga()

	filter := bson.M{configs.NAME: path.Base(r.URL.Path)}

	// Updates manga in DB
	if statusCode, err := updateEntity(configs.MANGAS_COLS, filter, configs.OP_SET,
		upd.Manga, upd.Search, upd.Filter); statusCode != 0 {
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

// DeleteManga .. Deletes manga from DB
func DeleteManga(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	// Deletes manga from DB
	if statusCode, err := deleteEntity(configs.MANGAS_COLS, bson.M{configs.NAME: path.Base(r.URL.Path)}); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
