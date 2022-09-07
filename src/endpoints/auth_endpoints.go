package endpoints

import (
	"net/http"

	"github.com/aivot-digital/preppy-backend/src/db"
	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/aivot-digital/preppy-backend/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := utils.ParseBody[models.Credentials](r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Bad request - Failed to parse body")
		return
	}

	user, err := db.Users.Get(bson.M{"username": credentials.Username})
	if err != nil {
		panic(err)
	}

	if user == nil || user.IsDeleted {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized - Invalid username or password")
		return
	}

	if !utils.CheckPassword(credentials.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized - Invalid username or password")
		return
	}

	accessToken, err := utils.CreateJwt(user, false)
	if err != nil {
		panic(err)
	}
	refreshToken, err := utils.CreateJwt(user, true)
	if err != nil {
		panic(err)
	}

	utils.SendJson(models.Jwt{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.Id,
	}, w)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Id: primitive.NewObjectID(),
	}

	// TODO

	accessToken, err := utils.CreateJwt(&user, false)
	if err != nil {
		panic(err)
	}
	refreshToken, err := utils.CreateJwt(&user, true)
	if err != nil {
		panic(err)
	}

	utils.SendJson(models.Jwt{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.Id,
	}, w)
}
