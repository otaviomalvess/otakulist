package configs

//=====================================
// Custom Validations
//=====================================
const (

	// content-status
	VAL_CONT_STATUS = "content-status"

	// content-type
	VAL_CONT_TYPE = "content-type"

	// content-producer
	VAL_CONT_PRODUCER = "content-producer"

	// content-studio
	VAL_CONT_STUDIO = "content-studio"

	// content-source
	VAL_CONT_SOURCE = "content-source"

	// content-genres
	VAL_CONT_GENRES = "content-genres"

	// content-anime-rating
	VAL_CONT_ANIME_RATE = "content-anime-rating"

	// content-book-rating
	VAL_CONT_BOOK_RATE = "content-book-rating"

	// content-season
	VAL_CONT_SEASON = "content-season"

	// content-season
	VAL_CONT_YEAR = "content-year"
)

//=====================================
// Maps for Validation
//=====================================
var (

	// Status
	VAL_STATUS = map[string]string{"Airing": "", "Not Yet Aired": "", "Finished": "", "Publishing": ""}

	// Type
	VAL_TYPE = map[string]string{"TV": "", "ONA": "", "OVA": "", "Manga": "", "Novel": ""}

	// Source
	VAL_SOURCE = map[string]string{"Manga": "", "Web Manga": "", "Novel": "", "Original": ""}

	// Anime Rating
	// Validates the rating according to   rating system
	VAL_ANIME_RATING = map[string]string{"G": "", "PG": "", "PG-13": "", "R": "", "NC-17": ""}

	// Manga Rating
	// Validates the rating according to Marvel rating system
	VAL_MANGA_RATING = map[string]string{"ALL AGES": "", "T": "", "T+ TEENS AND UP": "",
		"PARENTAL ADVISORY": "", "EXPLICIT CONTENT": ""}

	// Seasons
	VAL_SEASONS = map[string]string{"Spring": "", "Summer": "", "Fall": "", "Winter": ""}
)
