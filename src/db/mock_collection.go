package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockCollection[T any] struct {
	data []T
}

func NewMockCollection[T any]() *MockCollection[T] {
	return &MockCollection[T]{
		data: make([]T, 0),
	}
}

func (m *MockCollection[T]) Exists(filter bson.M) (bool, error) {
	return true, nil
}

func (m *MockCollection[T]) List(filter bson.M, limit uint64, offset uint64) ([]T, error) {
	return m.ListAndSort(filter, limit, offset, "")
}

func (m *MockCollection[T]) ListAndSort(filter bson.M, limit uint64, offset uint64, sortField string) ([]T, error) {
	return m.data, nil
}

func (m *MockCollection[T]) Create(data T) (*T, error) {
	m.data = append(m.data, data)
	return &data, nil
}

func (m *MockCollection[T]) Get(filter bson.M) (*T, error) {
	return nil, nil
}

func (m *MockCollection[T]) GetById(id string) (*T, error) {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Get(bson.M{"_id": itemId})
}

func (m *MockCollection[T]) Update(filter bson.M, update T) (*T, error) {
	return &update, nil
}

func (m *MockCollection[T]) UpdateById(id string, update T) (*T, error) {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Update(bson.M{"_id": itemId}, update)
}

func (m *MockCollection[T]) Remove(filter bson.M) error {
	return nil
}

func (m *MockCollection[T]) RemoveById(id string) error {
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return m.Remove(bson.M{"_id": itemId})
}

func (m *MockCollection[T]) Count(filter bson.M) (uint64, error) {
	return uint64(len(m.data)), nil
}
