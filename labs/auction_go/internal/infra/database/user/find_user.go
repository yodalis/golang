package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/yodalis/golang/labs/auction_go/config/logger"
	"github.com/yodalis/golang/labs/auction_go/internal/entity/user_entity"
	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"` // BSON -> indica como vão ser os campos dentro do mongo db
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context, userID string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userID}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this ID = %s", userID), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("User not found with this ID = %s", userID))
		}
		logger.Error("Error trying to find user by userID", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by userID")
	}

	userEntity := &user_entity.User{
		ID:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
