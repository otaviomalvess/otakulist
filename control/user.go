package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ListUser .. Lists the requested user
func ListUser(w http.ResponseWriter, r *http.Request) {

	// Gets user by ID in DB
	bytes, statusCode, err := listEntity(path.Base(r.URL.Path), configs.COL_USERS)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// SearchUsers .. Searches users by name
func SearchUsers(w http.ResponseWriter, r *http.Request) {

	// Searches for users by name in DB
	bytes, statusCode, err := searchEntities(r.URL.Query().Get(configs.NAME), configs.COL_SEARCH_USERS, configs.LIM_SEARCH_USERS)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}

// RegisterUser .. Registers the given user in the DB
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user model.RegistrableUser

	if err := util.BodyToModel(r.Body, &user.User); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    configs.TOKEN,
		Value:   user.User.Email,
		Expires: expire,
	}

	user.User.GenID = cookie.Value

	// Checks if anime already exists in DB
	if b, _ := checkExistence(user.User.Name, configs.COL_USERS); b {
		w.WriteHeader(http.StatusUnauthorized)
		util.WriteResponse(w, []byte(configs.ERR_EXIST_ENTITY))
		return
	}

	user.PopulateUserStructs()

	// Inserts user in DB
	if statusCode, err := registerEntity(configs.USERS_COLS, user.User, user.Search, user.AnimeList, user.MangaList); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	http.SetCookie(w, &cookie)
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// UpdateUser .. Updates user info
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var upd model.UpdatableUser

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &upd.User); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_BAD_REQ))
		return
	}

	upd.User.GenID = c.Value
	upd.PopulateUpdateUserStructs()

	// Updates user in DB
	if statusCode, err := updateEntity(configs.USERS_COLS, bson.M{configs.GEN_ID: upd.User.GenID}, configs.OP_SET,
		upd.User, upd.Search, upd.UpdatableAnimeList, upd.UpdatableMangaList); statusCode != 0 {

		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_UPDATING_ENTITY))
}

// DeleteUser .. Deletes user in DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	c, err := r.Cookie(configs.TOKEN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		return
	}

	// Deletes anime from DB
	if statusCode, err := deleteEntity(configs.USERS_COLS, bson.M{configs.GEN_ID: c.Value}); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
