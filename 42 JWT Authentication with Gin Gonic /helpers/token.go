package helpers

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fbdaf/go-jwt-gin/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func init() {
	if SECRET_KEY == "" {
		panic("SECRET_KEY is not set")
	}
}

func GenerateAllTokens(email, firstName, lastName, userType, userId string) (signedToken, signedRefreshToken string) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       userId,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 100).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}

	return token, refreshToken
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
	log.Printf("Attempting to validate token: %s", signedToken)

	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			log.Printf("Parsing token with claims")
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		log.Printf("Failed to cast claims")
		return nil, errors.New("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Printf("Token has expired")
		return nil, errors.New("token is expired")
	}

	log.Printf("Token validated successfully")
	return claims, nil
}

func UpdateAllTokens(token, refreshToken, userId string) error {
	var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: token})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: refreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(c, filter, bson.M{"$set": updateObj}, &opt)
	// _, err := userCollection.UpdateOne(c, filter, bson.D{"$set": updateObj}, &opt)
	return err
}
