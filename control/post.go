package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"path"

	"go.mongodb.org/mongo-driver/bson"
)

// Post .. Create new post in forum
func Post(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post model.Post

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filter := bson.M{configs.NAME: path.Base(r.URL.Path)}

	// Updates anime in DB
	if statusCode, err := updateEntityList(configs.COL_FORUMS, filter, configs.OP_PUSH,
		bson.M{configs.POSTS: post}, post); statusCode != 0 {

		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}

// EditPost .. Edits forum post
func EditPost(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post model.Post

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filter := bson.M{configs.NAME: path.Base(r.URL.Path), configs.OP_POSTS_USER_NAME: post.User.Name}
	upd := bson.M{configs.OP_POST_UPDATE: post.Post}

	// Updates anime in DB
	if statusCode, err := updateEntityList(configs.COL_FORUMS, filter, configs.OP_SET, upd, post); statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_UPDATING_ENTITY))
}

// DeletePost .. Deletes forum post
func DeletePost(w http.ResponseWriter, r *http.Request) {

	// Gets the cookie in request
	_, err := r.Cookie(configs.TOKEN)
	if err != nil {
		util.WriteResponse(w, []byte(configs.ERR_NOT_EXIST_COOKIE))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post model.Post

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Updates anime in DB
	if statusCode, err := updateEntityList(configs.COL_FORUMS, bson.M{}, configs.OP_PULL,
		bson.M{configs.POSTS: post}, post); statusCode != 0 {

		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, []byte(configs.SUCCESS_DELETING_ENTITY))
}
