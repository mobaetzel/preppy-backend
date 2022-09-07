package main

import (
	"net/http"

	"github.com/aivot-digital/preppy-backend/src/cfg"
	"github.com/aivot-digital/preppy-backend/src/db"
	"github.com/aivot-digital/preppy-backend/src/endpoints"
	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/aivot-digital/preppy-backend/src/repositories"
	"github.com/aivot-digital/preppy-backend/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	cfg.InitCfg()
	closeDb := db.InitDb()
	defer closeDb()

	prepareAdmin()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	recipeRepo := repositories.NewRecipeRepository(db.Recipes)
	userRepo := repositories.NewUserRepository(db.Users)

	router.Post("/v1/auth/login/", endpoints.Login)
	router.Get("/v1/auth/refresh/", endpoints.Refresh)

	router.Get("/v1/recipes/", endpoints.ListRecipes(recipeRepo))
	router.Post("/v1/recipes/", endpoints.CreateRecipe(recipeRepo))
	router.Get("/v1/recipes/{recipeId}/", endpoints.RetrieveRecipe(recipeRepo))
	router.Put("/v1/recipes/{recipeId}/", endpoints.UpdateRecipe(recipeRepo))
	router.Delete("/v1/recipes/{recipeId}/", endpoints.DeleteRecipe(recipeRepo))

	router.Get("/v1/users/", endpoints.ListUsers(userRepo))
	router.Post("/v1/users/", endpoints.CreateUser(userRepo))
	router.Get("/v1/users/{userId}/", endpoints.RetrieveUser(userRepo))
	router.Put("/v1/users/{userId}/", endpoints.UpdateUser(userRepo))
	router.Delete("/v1/users/{userId}/", endpoints.DeleteUser(userRepo))

	http.ListenAndServe("0.0.0.0:8000", router)
}

func prepareAdmin() {
	adminExists, err := db.Users.Exists(bson.M{"IsAdmin": true})
	if err != nil {
		panic(err)
	}
	if !adminExists {
		adminUsername := "admin"
		rawPassword := utils.RandomString(8)
		_, err := db.Users.Create(models.User{
			Id:        primitive.NewObjectID(),
			Username:  adminUsername,
			Password:  utils.HashPassword(rawPassword),
			IsAdmin:   true,
			IsDeleted: false,
		})
		if err != nil {
			panic(err)
		}
	}
}
