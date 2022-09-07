package repositories

import (
	"time"

	"github.com/aivot-digital/preppy-backend/src/db"
	"github.com/aivot-digital/preppy-backend/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeRepository struct {
	collection db.ICollection[models.Recipe]
}

func NewRecipeRepository(col db.ICollection[models.Recipe]) RecipeRepository {
	return RecipeRepository{
		collection: col,
	}
}

func (r *RecipeRepository) List(filter bson.M, pageSize uint64, pageIndex uint64) models.ListResponse[models.Recipe] {
	results, err := r.collection.ListAndSort(filter, pageSize, pageSize*pageIndex, "updated")
	if err != nil {
		panic(err)
	}
	count, err := r.collection.Count(filter)
	if err != nil {
		panic(err)
	}
	return models.ListResponse[models.Recipe]{
		Total: count,
		Items: results,
	}
}

func (r *RecipeRepository) Create(data models.Recipe) *models.Recipe {
	data.Id = primitive.NewObjectID()
	data.Created = time.Now()
	data.Updated = time.Now()

	result, err := r.collection.Create(data)
	if err != nil {
		panic(err)
	}

	return result
}

func (r *RecipeRepository) Get(id string) *models.Recipe {
	result, err := r.collection.GetById(id)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *RecipeRepository) Update(id string, update models.Recipe) *models.Recipe {
	update.Updated = time.Now()
	result, err := r.collection.UpdateById(id, update)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *RecipeRepository) Remove(id string) {
	err := r.collection.RemoveById(id)
	if err != nil {
		panic(err)
	}
}
