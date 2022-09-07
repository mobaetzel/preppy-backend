package db

import "go.mongodb.org/mongo-driver/bson"

type ICollection[T any] interface {
	Exists(filter bson.M) (bool, error)
	List(filter bson.M, limit uint64, offset uint64) ([]T, error)
	ListAndSort(filter bson.M, limit uint64, offset uint64, sortField string) ([]T, error)
	Create(data T) (*T, error)
	Get(filter bson.M) (*T, error)
	GetById(id string) (*T, error)
	Update(filter bson.M, update T) (*T, error)
	UpdateById(id string, update T) (*T, error)
	Remove(filter bson.M) error
	RemoveById(id string) error
	Count(filter bson.M) (uint64, error)
}
