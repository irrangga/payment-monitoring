package auth

import (
	"fmt"
	"net/http"
	"os"
	"payment-monitoring/models"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(authD models.Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUUID
	claims["username"] = authD.Username
	claims["role_id"] = authD.RoleID
	claims["role_name"] = authD.RoleName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*models.Auth, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string)
		if !ok {
			return nil, err
		}
		username, done := claims["username"].(string)
		if !done {
			return nil, err
		}
		roleId, done := claims["role_id"].(string)
		if !done {
			return nil, err
		}
		roleName, done := claims["role_name"].(string)
		if !done {
			return nil, err
		}
		return &models.Auth{
			AuthUUID: authUuid,
			Username: username,
			RoleID:   roleId,
			RoleName: roleName,
		}, nil
	}
	return nil, err
}
