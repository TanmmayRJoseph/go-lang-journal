package users

import (
	"context"
	"errors"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// created a collection of users
type UserRepo struct {
	col *mongo.Collection
}

// this function takes a database and returns a pointer to a UserRepo
func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{col: db.Collection("users")}

}

// find by email
func (r *UserRepo) CreateUser(ctx context.Context, u Users) (Users, error) {
	res, err := r.col.InsertOne(ctx, u)

	if err != nil {
		slog.Error("Error creating user", "error", err)
		return Users{}, err
	}

	u.Id = res.InsertedID.(bson.ObjectID)

	return u, nil
}

func (r *UserRepo) FindUserByEmail(ctx context.Context, email string) (Users, error) {
	filter := bson.M{"email": email}

	var user Users

	err := r.col.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			slog.Info("User not found", "email", email)
			return Users{}, mongo.ErrNoDocuments
		}

		slog.Error("Error finding user", "error", err)
		return Users{}, err
	}

	return user, nil
}