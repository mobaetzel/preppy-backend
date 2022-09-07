package endpoints

import (
	"net/http"

	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/aivot-digital/preppy-backend/src/repositories"
	"github.com/aivot-digital/preppy-backend/src/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func ListUsers(repo repositories.UserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, page := utils.GetPagination(r.URL.Query())
		filter := utils.GetFilter(r.URL.Query(), map[string]func(val string) (string, any){
			"search": func(val string) (string, any) {
				return "username", bson.M{"$text": val}
			},
		})
		Users := repo.List(filter, limit, page)
		utils.SendJson(Users, w)
	}
}

func CreateUser(repo repositories.UserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newUser, err := utils.ParseBody[models.User](r)
		if err != nil {
			utils.WriteBadRequest(w)
			return
		}

		User := repo.Create(newUser)

		err = utils.SendJson(User, w)
		if err != nil {
			panic(err)
		}
	}
}

func RetrieveUser(repo repositories.UserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		User := repo.Get(id)
		if User == nil {
			utils.WriteNotFound(w)
			return
		}

		err := utils.SendJson(User, w)
		if err != nil {
			panic(err)
		}
	}
}

func UpdateUser(repo repositories.UserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		updatedUser, err := utils.ParseBody[models.User](r)
		if err != nil {
			utils.WriteBadRequest(w)
			return
		}

		User := repo.Update(id, updatedUser)

		if User == nil {
			utils.WriteNotFound(w)
			return
		}

		err = utils.SendJson(User, w)
		if err != nil {
			panic(err)
		}
	}
}

func DeleteUser(repo repositories.UserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		repo.Remove(id)
		w.WriteHeader(http.StatusNoContent)
	}
}
