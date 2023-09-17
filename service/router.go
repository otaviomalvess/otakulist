package service

import (
	"net/http"
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/control"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var handler *mux.Router

// createCORS .. Creates the CORS to deal with cross origin
func createCORS() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodDelete},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})
}

// createHandler .. Creates the server router using the mux library
func createHandler() *mux.Router {

	// Creates a new router
	handler = mux.NewRouter()

	// Associates functions to paths
	// Index
	animeNews()

	// User
	user()
	login()
	profile()
	animeListProfile()
	animeListProfileOperations()
	mangaListProfile()
	mangaListProfileOperations()

	// Anime
	anime()
	animeProfile()

	//Manga
	manga()
	mangaProfile()

	// Forum
	forum()
	forumProfile()

	return handler
}

// Endpoint: /animeNews
func animeNews() {

	handler.HandleFunc(configs.ANIME_NEWS_PATH, control.AnimeNews).Methods(http.MethodGet).Queries(configs.SEASON, "", configs.YEAR, "")
}

// Endpoint: /user
func user() {

	handler.HandleFunc(configs.USER_PATH, control.SearchUsers).Methods(http.MethodGet).Queries(configs.NAME, "")
	handler.HandleFunc(configs.USER_PATH, control.RegisterUser).Methods(http.MethodPost).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /login
func login() {

	handler.HandleFunc(configs.LOGIN_PATH, control.Login).Methods(http.MethodPost)
}

// Endpoint: /profile/{username}
func profile() {

	handler.HandleFunc(configs.PROFILE_PATH, control.ListUser).Methods(http.MethodGet)
	handler.HandleFunc(configs.PROFILE_PATH, control.UpdateUser).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.PROFILE_PATH, control.DeleteUser).Methods(http.MethodDelete)
}

// Endpoint: /animelist/{username}
func animeListProfile() {

	handler.HandleFunc(configs.ANIME_LIST_PATH, control.AnimeList).Methods(http.MethodGet)
	handler.HandleFunc(configs.ANIME_LIST_PATH, control.AddAnimeInList).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /animelist/{username}/{anime}
func animeListProfileOperations() {

	handler.HandleFunc(configs.ANIME_LIST_OPERATIONS_PATH, control.UpdateAnimeInList).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.ANIME_LIST_OPERATIONS_PATH, control.DeleteAnimeInList).Methods(http.MethodDelete).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /mangalist/{username}
func mangaListProfile() {

	handler.HandleFunc(configs.MANGA_LIST_PATH, control.MangaList).Methods(http.MethodGet)
	handler.HandleFunc(configs.MANGA_LIST_PATH, control.AddMangaInList).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /mangalist/{username}/{manga}
func mangaListProfileOperations() {

	handler.HandleFunc(configs.MANGA_LIST_OPERATIONS_PATH, control.UpdateMangaInList).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.MANGA_LIST_OPERATIONS_PATH, control.DeleteMangaInList).Methods(http.MethodDelete).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /anime
func anime() {

	handler.HandleFunc(configs.ANIME_PATH, control.SearchAnimes).Methods(http.MethodGet).Queries(configs.ANIME, "")
	handler.HandleFunc(configs.ANIME_PATH, control.FilterAnimes).Methods(http.MethodGet).Queries(configs.NAME, "", configs.TYPE, "")
	handler.HandleFunc(configs.ANIME_PATH, control.RegisterAnime).Methods(http.MethodPost).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /anime/{anime}
func animeProfile() {

	handler.HandleFunc(configs.ANIME_PROFILE_PATH, control.ListAnime).Methods(http.MethodGet)
	handler.HandleFunc(configs.ANIME_PROFILE_PATH, control.UpdateAnime).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.ANIME_PROFILE_PATH, control.DeleteAnime).Methods(http.MethodDelete)
}

// Endpoint: /manga
func manga() {

	handler.HandleFunc(configs.MANGA_PATH, control.SearchMangas).Methods(http.MethodGet).Queries(configs.MANGA, "")
	handler.HandleFunc(configs.MANGA_PATH, control.FilterMangas).Methods(http.MethodGet).Queries(configs.NAME, "", configs.TYPE, "")
	handler.HandleFunc(configs.MANGA_PATH, control.RegisterManga).Methods(http.MethodPost).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /manga/{manga}
func mangaProfile() {

	handler.HandleFunc(configs.MANGA_PROFILE_PATH, control.ListManga).Methods(http.MethodGet)
	handler.HandleFunc(configs.MANGA_PROFILE_PATH, control.UpdateManga).Methods(http.MethodPut).Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.MANGA_PROFILE_PATH, control.DeleteManga).Methods(http.MethodDelete)
}

// Endpoint: /forum
func forum() {

	handler.HandleFunc(configs.FORUM_PATH, control.RegisterForum).Methods(http.MethodPost).Headers(configs.CONT_TYPE, configs.APP_JSON)
}

// Endpoint: /forum/{topic}
func forumProfile() {

	// Forum
	handler.HandleFunc(configs.FORUM_PROFILE_PATH, control.ListForum).Methods(http.MethodGet)
	handler.HandleFunc(configs.FORUM_PROFILE_PATH, control.DeleteForum).Methods(http.MethodDelete).Queries(configs.QRY_ACTION, "df")

	// Post
	handler.HandleFunc(configs.FORUM_PROFILE_PATH, control.Post).Methods(http.MethodPost).Queries(configs.QRY_ACTION, "post").Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.FORUM_PROFILE_PATH, control.EditPost).Methods(http.MethodPut).Queries(configs.QRY_ACTION, "edit").Headers(configs.CONT_TYPE, configs.APP_JSON)
	handler.HandleFunc(configs.FORUM_PROFILE_PATH, control.DeletePost).Methods(http.MethodDelete).Queries(configs.QRY_ACTION, "delete").Headers(configs.CONT_TYPE, configs.APP_JSON)
}
