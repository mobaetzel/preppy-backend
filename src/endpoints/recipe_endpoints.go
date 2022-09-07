package endpoints

import (
	"net/http"

	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/aivot-digital/preppy-backend/src/repositories"
	"github.com/aivot-digital/preppy-backend/src/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListRecipes(repo repositories.RecipeRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, page := utils.GetPagination(r.URL.Query())
		filter := utils.GetFilter(r.URL.Query(), map[string]func(val string) (string, any){
			"search": func(val string) (string, any) {
				return "title", bson.M{"$text": val}
			},
			"author-id": func(val string) (string, any) {
				authorId, err := primitive.ObjectIDFromHex(val)
				if err != nil {
					authorId = primitive.NewObjectID()
				}
				return "authorId", authorId
			},
			"tag": func(val string) (string, any) {
				return "tags", bson.M{"$elemMatch": val}
			},
		})
		recipes := repo.List(filter, limit, page)
		utils.SendJson(recipes, w)
	}
}

func CreateRecipe(repo repositories.RecipeRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newRecipe, err := utils.ParseBody[models.Recipe](r)
		if err != nil {
			utils.WriteBadRequest(w)
			return
		}

		recipe := repo.Create(newRecipe)

		err = utils.SendJson(recipe, w)
		if err != nil {
			panic(err)
		}
	}
}

func RetrieveRecipe(repo repositories.RecipeRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "recipeId")

		recipe := repo.Get(id)
		if recipe == nil {
			utils.WriteNotFound(w)
			return
		}

		err := utils.SendJson(recipe, w)
		if err != nil {
			panic(err)
		}
	}
}

func UpdateRecipe(repo repositories.RecipeRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "recipeId")
		updatedRecipe, err := utils.ParseBody[models.Recipe](r)
		if err != nil {
			utils.WriteBadRequest(w)
			return
		}

		recipe := repo.Update(id, updatedRecipe)

		if recipe == nil {
			utils.WriteNotFound(w)
			return
		}

		err = utils.SendJson(recipe, w)
		if err != nil {
			panic(err)
		}
	}
}

func DeleteRecipe(repo repositories.RecipeRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "recipeId")
		repo.Remove(id)
		w.WriteHeader(http.StatusNoContent)
	}
}
