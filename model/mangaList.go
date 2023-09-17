package model

// MangaList ..
type MangaList struct {
	GenID string        `bson:"genID" json:"genID" validate:"required"`
	Name  string        `bson:"name"  json:"name"  validate:"required"`
	List  []MangaInList `bson:"list"  json:"list"  validate:"required"`
}

// MangaInList .. Struct for manga info in users manga list
type MangaInList struct {
	Name             string           `bson:"name"      json:"name"      validate:"required"`
	Type             string           `bson:"type"      json:"type"      validate:"required,content-type"`
	ContentValuation ContentValuation `bson:"valuation" json:"valuation" validate:"required"`
}

// UpdatableMangaList ..
type UpdatableMangaList struct {
	GenID string `bson:"genID" json:"genID" validate:"required"`
	Name  string `bson:"name"  json:"name"  validate:"required"`
}

// UpdatableMangaInList .. Struct to update manga info in users manga list
type UpdatableMangaInList struct {
	Name string `bson:"name" json:"name" validate:"required"`
	Type string `bson:"type" json:"type" validate:"required,content-type"`
}
