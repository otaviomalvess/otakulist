package model

// Forum ..
type Forum struct {
	Name  string `bson:"name"  json:"name"  validate:"required"`
	Posts []Post `bson:"posts" json:"posts" validate:"required"`
}

// Post .. Forums post struct
type Post struct {
	User ForumUser `bson:"user" json:"user" validate:"required"`
	Post string    `bson:"post" json:"post" validate:"required"`
	Time int64     `bson:"time" json:"time" validate:"required"`
}
