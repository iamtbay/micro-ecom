package main

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
}

func initService() *Service {
	return &Service{}
}

var repo = initRepository()

//

func (x *Service) newFavorite(userID uuid.UUID, productID string) error {
	productObjID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}
	favoriteProductBSON := FavoriteProductBSON{
		ProductID: productObjID,
	}
	err = repo.newFavorite(userID, favoriteProductBSON)
	if err != nil {
		return err
	}
	return nil
}

func (x *Service) getAllFavorites(userID uuid.UUID) (AllFavorites, error) {
	favorites, err := repo.getAllFavorites(userID)
	if err != nil {
		return AllFavorites{}, err
	}
	return favorites, nil
}

func (x *Service) removeFavorite(userID uuid.UUID, productIDStr string) error {
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		return err
	}
	err = repo.removeFavorite(userID, productID)
	if err != nil {
		return err
	}
	return nil
}
