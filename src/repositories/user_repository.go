package repositories

import (
	"github.com/aivot-digital/preppy-backend/src/db"
	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/aivot-digital/preppy-backend/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	collection db.ICollection[models.User]
}

func NewUserRepository(col db.ICollection[models.User]) UserRepository {
	return UserRepository{
		collection: col,
	}
}

func (r *UserRepository) List(filter bson.M, pageSize uint64, pageIndex uint64) models.ListResponse[models.User] {
	results, err := r.collection.ListAndSort(filter, pageSize, pageSize*pageIndex, "_id")
	if err != nil {
		panic(err)
	}
	count, err := r.collection.Count(filter)
	if err != nil {
		panic(err)
	}
	for _, user := range results {
		user.Password = ""
	}
	return models.ListResponse[models.User]{
		Total: count,
		Items: results,
	}
}

func (r *UserRepository) Create(data models.User) *models.User {
	data.Id = primitive.NewObjectID()
	data.IsDeleted = false
	data.IsAdmin = false
	data.Password = utils.HashPassword(data.Password)

	result, err := r.collection.Create(data)
	if err != nil {
		panic(err)
	}

	return result
}

func (r *UserRepository) Get(id string) *models.User {
	result, err := r.collection.GetById(id)
	if err != nil {
		panic(err)
	}
	result.Password = ""
	return result
}

func (r *UserRepository) Update(id string, update models.User) *models.User {
	if update.Password != "" {
		update.Password = utils.HashPassword(update.Password)
	} else {
		user, err := r.collection.GetById(id)
		if err != nil {
			panic(err)
		}
		update.Password = user.Password
	}

	result, err := r.collection.UpdateById(id, update)
	if err != nil {
		panic(err)
	}

	result.Password = ""

	return result
}

func (r *UserRepository) Remove(id string) {
	user, err := r.collection.GetById(id)
	if err != nil {
		panic(err)
	}

	if user == nil {
		return
	}

	user.Username = ""
	user.Password = ""
	user.IsAdmin = false
	user.IsDeleted = true

	_, err = r.collection.UpdateById(id, *user)
	if err != nil {
		panic(err)
	}
}
