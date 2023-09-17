package model

// Anime ..
type Anime struct {
	Name      string   `bson:"name"      json:"name"       validate:"required"`
	Episodes  int16    `bson:"episodes"  json:"episodes"   validate:"numeric"`
	Producers []string `bson:"producers" json:"producers"  validate:"dive,content-producer"`
	Licensors string   `bson:"licensors" json:"licensors"`
	Studios   []string `bson:"studios"   json:"studios"    validate:"dive,content-studio"`
	Source    string   `bson:"source"    json:"source"     validate:"content-source"`
	Rating    string   `bson:"rating"    json:"rating"     validate:"content-anime-rating"`
	Season    string   `bson:"season"    json:"season"     validate:"content-season"`
	Year      int      `bson:"year"      json:"year"       validate:"numeric,min=1960,content-year"`

	*Content // Composes Anime type with Content type
}

// RegistrableAnime .. Anime struct for registration
type RegistrableAnime struct {
	Anime  Anime             `validate:"required"`
	Search SearchableContent `validate:"required"`
	Filter FilterableContent `validate:"required"`
}

// PopulateRegistrableAnime .. Populates the Registrable Anime struct
func (r *RegistrableAnime) PopulateRegistrableAnime() {

	r.Search = SearchableContent{
		Name:   r.Anime.Name,
		Status: r.Anime.Status,
	}

	r.Filter = FilterableContent{
		Name:      r.Anime.Name,
		Type:      r.Anime.Type,
		Status:    r.Anime.Status,
		Producers: r.Anime.Producers,
		Rating:    r.Anime.Rating,
		Genres:    r.Anime.Genres,
	}
}

// UpdatableAnime .. Anime struct for update
type UpdatableAnime struct {
	Anime  Anime                `validate:"required"`
	Search SearchableContent    `validate:"required"`
	Filter FilterableContent    `validate:"required"`
	InList UpdatableAnimeInList `validate:"required"`
}

// PopulateUpdatableAnime .. Populates the Updatable Anime struct
func (u *UpdatableAnime) PopulateUpdatableAnime() {

	u.Search = SearchableContent{
		Name:   u.Anime.Name,
		Status: u.Anime.Status,
	}

	u.Filter = FilterableContent{
		Name:      u.Anime.Name,
		Type:      u.Anime.Type,
		Status:    u.Anime.Status,
		Producers: u.Anime.Producers,
		Rating:    u.Anime.Rating,
		Genres:    u.Anime.Genres,
	}

	u.InList = UpdatableAnimeInList{
		Name: u.Anime.Name,
		Type: u.Anime.Type,
	}
}
