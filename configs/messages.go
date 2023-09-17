package configs

//=====================================
// Success
//=====================================
const (

	// Notifies success in connecting to DB
	SUCCESS_CONNECTING_DB = "[SUCCESS] Database connected"

	// Notifies successful entity registration to DB
	SUCCESS_REGISTRATING_ENTITY = `{"message": "Successful registration"}`

	// Notifies successful entity update in DB
	SUCCESS_UPDATING_ENTITY = `{"message": "Successful update"}`

	// Notifies successfull entity score update in DB
	SUCCESS_UPDATING_SCORE = `{"message": "Successful score update"}`

	// Notifies successful entity deletion of DB
	SUCCESS_DELETING_ENTITY = `{"message": "Successful deletion"}`
)

const (

	// Notifies error about not existing cookie
	ERR_NOT_EXIST_COOKIE = `{"message": "Not existent cookie"}`

	// Notifies error about already existent entity
	ERR_EXIST_ENTITY = `{"message": "Already existent entity"}`

	// Notifies internal server error
	ERR_INTERN_SERVER = `{"message": "Internal server error"}`

	// Notifies invalid login
	ERR_INVALID_LOGIN = `{"message": "Invalid login"}`

	// Notifies page not found
	ERR_NOT_FOUND = `{"message": "Not found"}`

	// Notifies bad request
	ERR_BAD_REQ = `{"message": "Bad request"}`
)

//=====================================
// Warnings
//=====================================
const (

	// Warns problem in parsing string
	WARN_PARSING_STR = "[WARN] problem parsing string, because, %v\n"

	// Warns problem in parsing JSON body
	WARN_PARSING_JSON = "[WARN] problem parsing JSON body, because, %v\n"

	// Warns problem in parsing JSON or BSON
	WARN_PARSING_SON = "[ERROR] parsing map to JSON or BSON, because, %v\n"

	// Warns problem in parsing io.ReadCloser body
	WARN_PARSING_IO = "[WARN] problem parsing io.ReadCloser body, because, %v\n"

	// Warns problem in connecting to DB
	WARN_CONNECTING_DB = "[WARN] Could not connect to database, because: %v\n"

	// Warns problem in validating struct
	WARN_VALIDATION = "[WARN] invalid struct information, because, %v\n"

	// Warns problem finding entity in DB
	WARN_FINDING_DOC = "[WARN] problem getting the document in database, because, %v\n"

	// Warns problem in decoding documents from DB
	WARN_DECODING_DOCS = "[WARN] problem decoding documents, because, %v\n"

	// Warns problem in inserting document in DB
	WARN_INSERTING_DOC = "[WARN] problem inserting document in database, because, %v\n"

	// Warns problem in updating document in DB
	WARN_UPDATING_DOC = "[WARN] problem updating document in database, because, %v\n"

	// Warns problem in deleting documento in DB
	WARN_DELETING_DOC = "[WARN] problem deleting document in database, because, %v\n"

	// Warns problem in fetching cookie
	WARN_COOKIE = "[WARN] error receiving cookie, because%v\n"
)

//=====================================
// Fatal
//=====================================
const (

	// Warns fatal error at creating client for DB
	FATAL_CLIENT = "[FATAL] Could not create client for database"

	// Warns fatal error at connectinf with DB
	FATAL_CONN_DB = "[FATAL] Could not connect to database"
)
