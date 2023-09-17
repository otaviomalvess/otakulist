package control

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/util"
)

// AnimeNews .. Gets seasonal anime
func AnimeNews(w http.ResponseWriter, r *http.Request) {

	b := util.MapToBSON(r.URL.Query())

	// Gets user by ID in DB
	bytes, statusCode, err := filterEntity(b, configs.COL_ANIMES)
	if statusCode != 0 {
		w.WriteHeader(statusCode)
		util.WriteResponse(w, []byte(err))
		return
	}

	// Calls write response util function
	util.WriteResponse(w, bytes)
}
