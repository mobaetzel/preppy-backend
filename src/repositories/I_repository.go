package repositories

import (
	"github.com/aivot-digital/preppy-backend/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IRepository[T any] interface {
	List(filter bson.M, pageSize uint64, pageIndex uint64) models.ListResponse[T]
	Create(data T) *T
	Get(id string) *T
	Update(id string, update T) *T
	Remove(id string)
}
