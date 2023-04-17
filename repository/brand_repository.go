package repository

import (
	"context"
	"github.com/jutionck/go-sinar-makmur-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BrandRepository interface {
	Create(payload *model.Brand) error
	List() ([]model.Brand, error)
	Get(id string) (model.Brand, error)
}

type brandRepository struct {
	db *mongo.Database
}

func (b *brandRepository) Get(id string) (model.Brand, error) {
	var brand model.Brand
	// Karena _id adalah sebuah objectID maka kita harus lakukan convert dulu dari string ke objectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Brand{}, err
	}
	err = b.db.Collection("brand").FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&brand)
	if err != nil {
		return model.Brand{}, err
	}
	log.Println("BrandRepository.Get:", brand)
	return brand, nil
}

func (b *brandRepository) Create(payload *model.Brand) error {
	payload.Id = primitive.NewObjectID() // untuk menambahkan _id dengan tipe data objectID
	brand, err := b.db.Collection("brand").InsertOne(context.Background(), payload)
	if err != nil {
		return err
	}
	log.Println("BrandRepository.Create:", brand)
	return nil
}

func (b *brandRepository) List() ([]model.Brand, error) {
	var brands []model.Brand
	// mengambil data dari collection "brand", bson.M{} -> kita mengambil seluruh dokumen yang ada pada collection
	cursor, err := b.db.Collection("brand").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var brand model.Brand
		// bson.M{} -> struct -> brand
		err := cursor.Decode(&brand)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	log.Println("BrandRepository.List:", brands)
	return brands, nil
}

func NewBrandRepository(db *mongo.Database) BrandRepository {
	return &brandRepository{db: db}
}
