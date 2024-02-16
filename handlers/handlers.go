package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"

	"auth-service/db"
	"auth-service/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetTokensHandler(w http.ResponseWriter, r *http.Request) {
	var tokenReq models.TokenRequest
	_ = json.NewDecoder(r.Body).Decode(&tokenReq)

	accessToken := jwt.New(jwt.SigningMethodHS512)
	accessToken.Claims = jwt.MapClaims{
		"userId": tokenReq.UserID,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}
	accessTokenString, _ := accessToken.SignedString([]byte("12345")) // секретный JWT токен

	refreshUuid := uuid.New()
	refreshUuidBytes, _ := refreshUuid.MarshalBinary()
	refreshHash, _ := bcrypt.GenerateFromPassword(refreshUuidBytes, bcrypt.DefaultCost)
	refreshHashString := base64.StdEncoding.EncodeToString(refreshHash)

	coll := db.GetCollection()
	_, err := coll.InsertOne(context.Background(), models.Token{RefreshTokenHash: refreshHashString})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.TokenDetails{
		AccessToken:  accessTokenString,
		RefreshToken: refreshHashString,
	})
}

func RefreshTokensHandler(w http.ResponseWriter, r *http.Request) {
	refreshHashString := r.Header.Get("Authorization")

	coll := db.GetCollection()
	var storedToken models.Token
	err := coll.FindOne(context.Background(), bson.M{"refreshTokenHash": refreshHashString}).Decode(&storedToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	refreshHash, _ := base64.StdEncoding.DecodeString(storedToken.RefreshTokenHash)

	err = bcrypt.CompareHashAndPassword(refreshHash, []byte(refreshHashString))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken := jwt.New(jwt.SigningMethodHS512)
	accessToken.Claims = jwt.MapClaims{
		"userId": "TODO",
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}
	accessTokenString, _ := accessToken.SignedString([]byte("55555")) // секретный JWT токен
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.TokenDetails{
		AccessToken:  accessTokenString,
		RefreshToken: refreshHashString,
	})
}
