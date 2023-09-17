package model

// AnimeList ..
type AnimeList struct {
	GenID string        `bson:"genID" json:"genID" validate:"required"`
	Name  string        `bson:"name"  json:"name"  validate:"required"`
	List  []AnimeInList `bson:"list"  json:"list"  validate:"required"`
}

// AnimeInList .. Struct for anime info in users anime list
type AnimeInList struct {
	Name             string           `bson:"name"      json:"name"      validate:"required"`
	Type             string           `bson:"type"      json:"type"      validate:"required,content-type"`
	ContentValuation ContentValuation `bson:"valuation" json:"valuation" validate:"required"`
}

// UpdatableAnimeList .. Struct to update user anime list info
type UpdatableAnimeList struct {
	GenID string `bson:"genID" json:"genID" validate:"required"`
	Name  string `bson:"name"  json:"name"  validate:"required"`
}

// UpdatableAnimeInList .. Struct to update anime info in users anime list
type UpdatableAnimeInList struct {
	Name string `bson:"name" json:"name" validate:"required"`
	Type string `bson:"type" json:"type" validate:"required,content-type"`
}
