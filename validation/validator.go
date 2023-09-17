package validation

import (
	"otavio-alves/OtakuList/configs"
	"otavio-alves/OtakuList/database"
	"time"

	val "gopkg.in/go-playground/validator.v9"
)

// Validator ..
var Validator *val.Validate

// CreateValidator .. Creates a new validator
func CreateValidator() {

	Validator = val.New()

	// Registers custom validations for structs
	Validator.RegisterValidation(configs.VAL_CONT_STATUS, ValidateStatus)
	Validator.RegisterValidation(configs.VAL_CONT_TYPE, ValidateType)
	Validator.RegisterValidation(configs.VAL_CONT_PRODUCER, ValidateProducer)
	Validator.RegisterValidation(configs.VAL_CONT_STUDIO, ValidateStudio)
	Validator.RegisterValidation(configs.VAL_CONT_SOURCE, ValidateSource)
	Validator.RegisterValidation(configs.VAL_CONT_GENRES, ValidateGenre)
	Validator.RegisterValidation(configs.VAL_CONT_ANIME_RATE, ValidateAnimeRating)
	Validator.RegisterValidation(configs.VAL_CONT_BOOK_RATE, ValidateBookRating)
	Validator.RegisterValidation(configs.VAL_CONT_SEASON, ValidateSeason)
	Validator.RegisterValidation(configs.VAL_CONT_YEAR, ValidateYear)
}

// ValidateStatus .. Validates the current status of the content
func ValidateStatus(fl val.FieldLevel) bool {

	status := fl.Field().String()
	_, ok := configs.VAL_STATUS[status]
	return ok
}

// ValidateType .. Validates the type of the content
func ValidateType(fl val.FieldLevel) bool {

	contType := fl.Field().String()
	_, ok := configs.VAL_TYPE[contType]
	return ok
}

// ValidateProducer .. Validates the producer of the anime
func ValidateProducer(fl val.FieldLevel) bool {

	producer := fl.Field().String()
	return database.FindForValidation(producer, configs.COL_PRODUCERS)
}

// ValidateStudio .. Validates the studio of the anime
func ValidateStudio(fl val.FieldLevel) bool {

	studio := fl.Field().String()
	return database.FindForValidation(studio, configs.COL_STUDIOS)
}

// ValidateSource .. Validates the original source of the anime
func ValidateSource(fl val.FieldLevel) bool {

	source := fl.Field().String()
	_, ok := configs.VAL_SOURCE[source]
	return ok
}

// ValidateGenre .. Validates the genre of the content
func ValidateGenre(fl val.FieldLevel) bool {

	genre := fl.Field().String()
	return database.FindForValidation(genre, configs.COL_GENRES)
}

// ValidateAnimeRating .. Validates the anime rating
func ValidateAnimeRating(fl val.FieldLevel) bool {

	rating := fl.Field().String()
	_, ok := configs.VAL_ANIME_RATING[rating]
	return ok
}

// ValidateBookRating .. Validates the manga / novel rating
func ValidateBookRating(fl val.FieldLevel) bool {

	rating := fl.Field().String()
	_, ok := configs.VAL_MANGA_RATING[rating]
	return ok
}

// ValidateSeason .. Validates the anime season
func ValidateSeason(fl val.FieldLevel) bool {

	season := fl.Field().String()
	_, ok := configs.VAL_SEASONS[season]
	return ok
}

// ValidateYear .. Validates the year of the anime season
func ValidateYear(fl val.FieldLevel) bool {

	year := int(fl.Field().Int())
	return year <= time.Now().Year()
}
