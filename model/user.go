package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ..
type User struct {

	// TODO: have an image as an user's profile picture

	GenID    string `bson:"genID"    json:"genID"    validate:"required"`
	Name     string `bson:"name"     json:"name"     validate:"required"`
	Email    string `bson:"email"    json:"email"    validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

// RegistrableUser .. User struct for registration
type RegistrableUser struct {
	User      User       `validate:"required"`
	Search    UserSearch `validate:"required"`
	AnimeList AnimeList  `validate:"required"`
	MangaList MangaList  `validate:"required"`
}

// PopulateUserStructs ..
func (u *RegistrableUser) PopulateUserStructs() {

	u.Search = UserSearch{
		GenID: u.User.GenID,
		Name:  u.User.Name,
	}

	u.AnimeList = AnimeList{
		GenID: u.User.GenID,
		Name:  u.User.Name,
		List:  []AnimeInList{},
	}

	u.MangaList = MangaList{
		GenID: u.User.GenID,
		Name:  u.User.Name,
		List:  []MangaInList{},
	}
}

// UpdatableUser .. User struct for update
type UpdatableUser struct {
	User               User               `validate:"required"`
	Search             UserSearch         `validate:"required"`
	UpdatableAnimeList UpdatableAnimeList `validate:"required"`
	UpdatableMangaList UpdatableMangaList `validate:"required"`
}

// PopulateUpdateUserStructs ..
func (u *UpdatableUser) PopulateUpdateUserStructs() {

	u.Search = UserSearch{
		GenID: u.User.GenID,
		Name:  u.User.Name,
	}

	u.UpdatableAnimeList = UpdatableAnimeList{
		GenID: u.User.GenID,
		Name:  u.User.Name,
	}

	u.UpdatableMangaList = UpdatableMangaList{
		GenID: u.User.GenID,
		Name:  u.User.Name,
	}
}

// UserSearch .. User struct for search
type UserSearch struct {

	// TODO: have an image as an user's profile picture

	GenID string `bson:"genID" json:"genID" validate:"required"`
	Name  string `bson:"name"  json:"name"  validate:"required"`
}

// ForumUser ..
type ForumUser struct {
	Name string `bson:"name" json:"name" validate:"required"`
}

// UserID .. TODO: eliminate this struct
type UserID struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
}
