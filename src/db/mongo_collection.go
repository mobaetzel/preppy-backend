package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection[T any] struct {
	collection *mongo.Collection
}

func NewMongoCollection[T any](db *mongo.Database, collectionName string) *MongoCollection[T] {
	return &MongoCollection[T]{
		collection: db.Collection(collectionName),
	}
}

func (m *MongoCollection[T]) Exists(filter bson.M) (bool, error) {
	ctx := context.TODO()

	result := m.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		} else {
			return false, result.Err()
		}
	}

	return true, nil
}
func (m *MongoCollection[T]) List(filter bson.M, limit uint64, offset uint64) ([]T, error) {
	return m.ListAndSort(filter, limit, offset, "")
}

func (m *MongoCollection[T]) ListAndSort(filter bson.M, limit uint64, offset uint64, sortField string) ([]T, error) {
	opts := options.
		Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit))

	if sortField != "" {
		opts.SetSort(bson.M{
			sortField: -1,
		})
	}

	cursor, err := m.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	results := make([]T, 0)
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var item T
		err = cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}

func (m *MongoCollection[T]) Create(data T) (*T, error) {
	_, err := m.collection.InsertOne(context.TODO(), data)
	return &data, err
}

func (m *MongoCollection[T]) Get(filter bson.M) (*T, error) {
	result := m.collection.FindOne(context.TODO(), filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, result.Err()
		}
	}

	var item T
	err := result.Decode(&item)

	return &item, err
}

func (m *MongoCollection[T]) GetById(id string) (*T, error) {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Get(bson.M{"_id": itemId})
}

func (m *MongoCollection[T]) Update(filter bson.M, update T) (*T, error) {
	res, err := m.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	if res.ModifiedCount == 0 {
		return nil, nil
	}

	return &update, nil
}

func (m *MongoCollection[T]) UpdateById(id string, update T) (*T, error) {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Update(bson.M{"_id": itemId}, update)
}

func (m *MongoCollection[T]) Remove(filter bson.M) error {
	_, err := m.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (m *MongoCollection[T]) RemoveById(id string) error {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return m.Remove(bson.M{"_id": itemId})
}

func (m *MongoCollection[T]) Count(filter bson.M) (uint64, error) {
	count, err := m.collection.CountDocuments(context.TODO(), filter)
	return uint64(count), err
}
