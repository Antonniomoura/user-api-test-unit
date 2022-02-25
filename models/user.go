package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	CreateAt time.Time          `bson:"create_at" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at" json:"update_at,omitempty"`
}

type Users []*User
