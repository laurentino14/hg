package utils

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"errors"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

type Subject struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
}

func ValidateAT(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, mt := token.Method.(*jwt.SigningMethodHMAC); !mt {

			return nil, errors.New("unexpected signing method")
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}

		subjectJSON := Subject{}
		err = json.Unmarshal([]byte(subject), &subjectJSON)
		if err != nil {
			return nil, err
		}

		var typeUser string

		switch subjectJSON.Type {
		case 0:
			typeUser = "CLIENT"
		case 1:
			typeUser = "ADMIN"
		case 2:
			typeUser = "WORKER"
		}

		secretKey := "JWT_SECRET_AT_" + typeUser

		secret := []byte(os.Getenv(secretKey))

		return secret, nil

	})

}

func ValidateRT(token string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, mt := token.Method.(*jwt.SigningMethodHMAC); !mt {
			return nil, errors.New("unexpected signing method")
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}

		subjectJSON := Subject{}
		err = json.Unmarshal([]byte(subject), &subjectJSON)
		if err != nil {
			return nil, err
		}

		var typeUser string

		switch subjectJSON.Type {
		case 0:
			typeUser = "CLIENT"
		case 1:
			typeUser = "ADMIN"
		case 2:
			typeUser = "WORKER"
		}

		secretKey := "JWT_SECRET_RT_" + typeUser

		secret := []byte(os.Getenv(secretKey))

		return secret, nil

	})

}

func GenerateAccessToken(user contracts.UserContract) (string, error) {
	var typeUser string
	switch user.Type {
	case 0:
		typeUser = "CLIENT"
	case 1:
		typeUser = "ADMIN"
	case 2:
		typeUser = "WORKER"
	}

	subject := Subject{
		ID:   user.ID,
		Type: int(user.Type),
	}

	subjectBYTES, err := json.Marshal(subject)

	if err != nil {
		return "", err
	}

	claimsAT := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(subjectBYTES),
		}}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAT)
	secretKey := "JWT_SECRET_AT_" + typeUser

	token, err := t.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(user contracts.UserContract) (string, error) {
	var typeUser string
	switch user.Type {
	case 0:
		typeUser = "CLIENT"
	case 1:
		typeUser = "ADMIN"
	case 2:
		typeUser = "WORKER"
	}

	subject := Subject{
		ID:   user.ID,
		Type: int(user.Type),
	}

	subjectBYTES, err := json.Marshal(subject)

	if err != nil {
		return "", err
	}

	claimsRT := &Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(subjectBYTES),
		}}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRT)

	secretKey := "JWT_SECRET_RT_" + typeUser

	token, err := t.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		return "", err
	}

	return token, nil
}
