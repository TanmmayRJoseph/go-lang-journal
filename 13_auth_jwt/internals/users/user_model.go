package users

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Users struct {
	Id           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Email        string        `bson:"email" json:"email"`
	PasswordHash string        `bson:"password" json:"password"`
	Role         string        `bson:"role" json:"role"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
}

type PublicUser struct {
	Id           string    ` json:"id"`
	Name         string    ` json:"name"`
	Email        string    ` json:"email"`
	PasswordHash string    ` json:"password"`
	Role         string    ` json:"role"`
	CreatedAt    time.Time ` json:"created_at"`
	UpdatedAt    time.Time ` json:"updated_at"`
}

func ToPublic(u Users) PublicUser {
	// converting a user to a public user
	return PublicUser{
		Id:           u.Id.Hex(), // converting ObjectID to string
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
