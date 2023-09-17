package model

// Content ..
type Content struct {

	// TODO: have an image for content cover

	Type            string   `bson:"type"             json:"type"              validate:"required,content-type"`
	Status          string   `bson:"status"           json:"status"            validate:"required,content-status"`
	LaunchingStart  string   `bson:"launching_start"  json:"launching_start"`
	LaunchingFinish string   `bson:"launching_finish" json:"launching_finish"`
	Genres          []string `bson:"genres"           json:"genres"            validate:"required,dive,required,content-genres"`
	Synopsis        string   `bson:"synopsis"         json:"synopsis"`
}

// SearchableContent .. Content struct for search
type SearchableContent struct {

	// TODO: have an image for content cover

	Name   string `bson:"name"   json:"name"   validate:"required"`
	Status string `bson:"status" json:"status" validate:"required,content-status"`
}

// FilterableContent .. Content struct for filtered search
type FilterableContent struct {
	Name      string   `bson:"name"     json:"name"`
	Type      string   `bson:"type"     json:"type"`
	Status    string   `bson:"status"   json:"status"`
	Producers []string `bson:"producer" json:"producer"`
	Rating    string   `bson:"rating"   json:"rating"`
	Genres    []string `bson:"genres"   json:"genres"`
}

// UserContentValuation ..
type UserContentValuation struct {
	Name  string `bson:"name"  json:"name"  validate:"required"`
	Score int    `bson:"score" json:"score" validate:"min=0,max=10"`
}

// ContentValuation .. Content struct for content valuation by user
type ContentValuation struct {
	Score    int `bson:"score"    json:"score"    validate:"min=0,max=10"`
	Progress int `bson:"progress" json:"progress" validate:"min=0"`
}
