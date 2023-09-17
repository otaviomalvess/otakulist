package configs

//=====================================
// Server HTTP
//=====================================
const (

	// Defines IP address and port for server instance
	SERVER_ADDRS = "localhost:8080"
)

//=====================================
// Paths
//=====================================
const (

	// Defines recommended path
	RECOMMENDED_PATH = "/recommended"

	// Defines anime news path
	ANIME_NEWS_PATH = "/animenews"

	// Defines user path
	USER_PATH = "/user"

	// Defines login path
	LOGIN_PATH = "/login"

	// Define profile path
	PROFILE_PATH = "/profile/{username}"

	// Defines anime list path
	ANIME_LIST_PATH = "/animelist/{username}"

	// Defines anime list operations path
	ANIME_LIST_OPERATIONS_PATH = "/animelist/{username}/{anime}"

	// Defines manga list path
	MANGA_LIST_PATH = "/mangalist/{username}"

	// Defines manga list operations path
	MANGA_LIST_OPERATIONS_PATH = "/mangalist/{username}/{manga}"

	// Defines anime path
	ANIME_PATH = "/anime"

	// Defines anime profile path
	ANIME_PROFILE_PATH = "/anime/{anime}"

	// Defines manga path
	MANGA_PATH = "/manga"

	// Defines manga profile path
	MANGA_PROFILE_PATH = "/manga/{manga}"

	// Defines anime score profile path
	ANIME_SCORE_PROFILE_PATH = "/score/anime/{anime}"

	// Defines manga score profile path
	MANGA_SCORE_PROFILE_PATH = "/score/manga/{manga}"

	// Defines forum path
	FORUM_PATH = "/forum"

	// Defines forum profile path
	FORUM_PROFILE_PATH = "/forum/{topic}"
)

//=====================================
// Queries
//=====================================
const (

	// Action
	QRY_ACTION = "action"
)

//=====================================
// Headers
//=====================================
const (

	// Content-Type
	CONT_TYPE = "Content-Type"

	// application/json
	APP_JSON = "application/json"
)
