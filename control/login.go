package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/model"
	"otavio-alves/OtakuList/util"
	"time"
)

// Login ..
func Login(w http.ResponseWriter, r *http.Request) {

	var login model.Login

	// Calls parsing util functions and checks if any error occurs during
	if err := util.BodyToModel(r.Body, &login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Gets user by ID in DB
	bytes, statusCode, err := loginUser(login.Email, configs.COL_USERS)
	if statusCode != 0 || bytes == nil {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	sha256 := login.Email
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    configs.TOKEN,
		Value:   string(sha256[:]),
		Expires: expire,
	}

	// Calls write response util function
	http.SetCookie(w, &cookie)
	util.WriteResponse(w, []byte(configs.SUCCESS_REGISTRATING_ENTITY))
}
