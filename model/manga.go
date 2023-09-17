package model

// Manga ..
type Manga struct {
	Name          string   `bson:"name"          json:"name"           validate:"required"`
	Volumes       int16    `bson:"volumes"       json:"volumes"        validate:"numeric"`
	Chapters      int16    `bson:"chapter"       json:"chapter"        validate:"numeric"`
	Authors       []string `bson:"authors"       json:"authors"        validate:"required"`
	Serialization string   `bson:"serialization" json:"serialization"`
	Rating        string   `bson:"rating"        json:"rating"         validate:"content-book-rating"`

	// Composes manga type with content type
	*Content
}

// RegistrableManga .. Manga struct for registration
type RegistrableManga struct {
	Manga  Manga             `validate:"required"`
	Search SearchableContent `validate:"required"`
	Filter FilterableContent `validate:"required"`
}

// PopulateRegistrableManga .. Populates the Registrable Manga struct
func (r *RegistrableManga) PopulateRegistrableManga() {

	r.Search = SearchableContent{
		Name:   r.Manga.Name,
		Status: r.Manga.Status,
	}

	r.Filter = FilterableContent{
		Name:      r.Manga.Name,
		Type:      r.Manga.Type,
		Status:    r.Manga.Status,
		Producers: r.Manga.Authors,
		Rating:    r.Manga.Rating,
		Genres:    r.Manga.Genres,
	}
}

// UpdatableManga .. Manga struct for update
type UpdatableManga struct {
	Manga  Manga                `validate:"required"`
	Search SearchableContent    `validate:"required"`
	Filter FilterableContent    `validate:"required"`
	InList UpdatableMangaInList `validate:"required"`
}

// PopulateUpdatableManga .. Populates the Updatable Manga struct
func (u *UpdatableManga) PopulateUpdatableManga() {

	u.Search = SearchableContent{
		Name:   u.Manga.Name,
		Status: u.Manga.Status,
	}

	u.Filter = FilterableContent{
		Name:      u.Manga.Name,
		Type:      u.Manga.Type,
		Status:    u.Manga.Status,
		Producers: u.Manga.Authors,
		Rating:    u.Manga.Rating,
		Genres:    u.Manga.Genres,
	}

	u.InList = UpdatableMangaInList{
		Name: u.Manga.Name,
		Type: u.Manga.Type,
	}
}
