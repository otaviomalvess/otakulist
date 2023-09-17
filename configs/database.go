package configs

//=====================================
// Database
//=====================================
const (

	// Defines the name of the database
	DB_NAME = "otakuListDB"

	// Defines IP address and port for database instance
	MONGO_HOST = "mongodb://localhost:27017"
)

//=====================================
// Collections
//=====================================
const (

	// Defines users collection
	COL_USERS = "users"

	// Defines animes collection
	COL_ANIMES = "animes"

	// Defines mangas collection
	COL_MANGAS = "mangas"

	// Defines genres collection
	COL_GENRES = "genres"

	// Defines producers collection
	COL_PRODUCERS = "producers"

	// Defines studios collection
	COL_STUDIOS = "studios"

	// Defines animes collection for search
	COL_SEARCH_ANIMES = "searchAnimes"

	// Defines mangas collection for search
	COL_SEARCH_MANGAS = "searchMangas"

	// Defines users collection for search
	COL_SEARCH_USERS = "searchUsers"

	// Defines animes collection for search
	COL_FILTER_ANIMES = "filterAnimes"

	// Defines mangas collection for search
	COL_FILTER_MANGAS = "filterMangas"

	// Defines collection for users anime lists
	COL_ANIME_LIST = "animeLists"

	// Defines collection for users manga lists
	COL_MANGA_LIST = "mangaLists"

	// Defines collection for forums
	COL_FORUMS = "forums"

	// Defines collection for anime news
	COL_NEWS = "news"
)

//=====================================
// Slices of Collections Names
//=====================================
var (

	// Defines animes collection array
	ANIMES_COLS = []string{COL_ANIMES, COL_SEARCH_ANIMES, COL_FILTER_ANIMES}

	// Defines mangas collection array
	MANGAS_COLS = []string{COL_MANGAS, COL_SEARCH_MANGAS, COL_FILTER_MANGAS}

	// Defines users collection array
	USERS_COLS = []string{COL_USERS, COL_SEARCH_USERS, COL_ANIME_LIST, COL_MANGA_LIST}
)

//=====================================
// Operators
//=====================================
const (

	// Set operator
	OP_SET = "$set"

	// Push operator
	OP_PUSH = "$push"

	// Pull operator
	OP_PULL = "$pull"

	// List name operator
	OP_LIST_NAME = "list.name"

	// List index operator
	OP_LIST_INDEX = "list.$"

	// List name update operator
	OP_LIST_NAME_UPDATE = "list.$.name"

	// List type update operator
	OP_LIST_TYPE_UPDATE = "list.$.type"

	// List valuation update operator
	OP_LIST_VALUATION_UPDATE = "list.$.valuation"

	// List content status update operator
	OP_CONTENT_STATUS_UPD = "content_status.users_valuation.$.name"

	// Post user name
	OP_POSTS_USER_NAME = "posts.user.name"

	// Post update operator
	OP_POST_UPDATE = "posts.$.post"
)
