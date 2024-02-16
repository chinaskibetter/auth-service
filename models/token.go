package models

type Token struct {
	RefreshTokenHash string `bson:"refreshTokenHash"`
}

type TokenRequest struct {
	UserID string `json:"userId"`
}

type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
