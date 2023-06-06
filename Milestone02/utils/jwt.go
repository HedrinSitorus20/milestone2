package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

func GenerateJwt(username string, roleid uint) (string, error) {
	claims := jwt.MapClaims{
		//"sub":  userId,
		"name":    username,
		"role_id": roleid,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var signedToken, err = token.SignedString([]byte("hedrin"))
	return signedToken, err
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hedrin"), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyJWT(token string) (bool, error) {
	hasil, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hedrin"), nil
	})

	if err != nil {
		return false, err
	}

	if a, ok := hasil.Claims.(jwt.MapClaims); ok && hasil.Valid {
		fmt.Println(a)
		return true, nil
	}
	return false, err
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "hedrin"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
