package authorization

import (
	"errors"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *entity.User) (string, error) {
	claim := entity.ClaimUser{
		Id:       user.Id,
		Email:    user.Email,
		Dni:      user.Dni,
		Name:     user.Name,
		LastName: user.LastName,
		Rol:      user.Rol,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 50).Unix(),
			Issuer:    "Leonardo Antonio Nolasco Leyva",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (entity.ClaimUser, error) {
	token, err := jwt.ParseWithClaims(t, &entity.ClaimUser{}, verifyFunction)
	if err != nil {
		return entity.ClaimUser{}, err
	}
	if !token.Valid {
		return entity.ClaimUser{}, errors.New("el token no es valido")
	}

	claim, ok := token.Claims.(*entity.ClaimUser)
	if !ok {
		return entity.ClaimUser{}, errors.New("no se pudieron obtener claims")
	}
	return *claim, nil
}

func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
