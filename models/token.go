package models

type Token struct {
	RefreshTokenHash string `bson:"refreshTokenHash"`
}
