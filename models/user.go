package models

import (
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 32
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Email       string             `json:"email"`
	Password    string             `json:"-"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	AccessToken AccessToken        `json:"accessToken" bson:"inline"`
}

type AccessToken struct {
	Value     string        `json:"value"`
	TTL       time.Duration `json:"ttl"`
	ExpiresAt time.Time     `json:"expiresAt"`
}

type RegisterUserBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (body *RegisterUserBody) ToUserDocument() (*User, *HttpRequestError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcryptCost)
	if err != nil {
		return nil, InternalServerError("Unable to hash the provided password")
	}

	return &User{
		Email:     body.Email,
		Password:  string(hashedPassword),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		AccessToken: AccessToken{
			Value:     RandomString(256),
			TTL:       3 * time.Hour,
			ExpiresAt: time.Now().Add(3 * time.Hour),
		},
	}, nil
}
