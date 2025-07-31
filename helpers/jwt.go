package helpers

import (
	"akatgelar/orderfaz-testcase-be/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyJWT(tokenString string, publicKeyPath string) (jwt.MapClaims, error) {
    // Read the public key from file
    publicKeyData, err := os.ReadFile(publicKeyPath)
    if err != nil {
        return nil, fmt.Errorf("could not read public key: %v", err)
    }

    // Parse the public key
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
    if err != nil {
        return nil, fmt.Errorf("could not parse public key: %v", err)
    }

    // Parse and verify the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return publicKey, nil
    })
    if err != nil {
        return nil, fmt.Errorf("could not parse token: %v", err)
    }

    claims, err := ExtractClaims(token) 
    if err != nil {
        return nil, fmt.Errorf("could not extract claim: %v", err)
    } 

    exp := claims["exp"].(float64)
    expirationTime := time.Unix(int64(exp), 0)  
    if expirationTime.Before(time.Now()) {
        return nil, fmt.Errorf("token has expired: %v", err) 
    }  

    return claims, nil
}

func ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("invalid token")
} 

func JwtMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {
        err := godotenv.Load()
        if err != nil { 
            fmt.Println(err)
        } 
        secretKey := []byte(os.Getenv("SECRET_KEY")) 
        
		var message string
		header := c.GetHeader("Authorization")
		header_split := strings.Split(header, " ") 

		if len(header_split) == 2 {
			if strings.ToLower(header_split[0]) == "bearer" {
				tokenString := header_split[1] 

                token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                    return secretKey, nil 
                }) 
                
                if err != nil {
                    message = "Internal error, " + err.Error()
                    response := models.BaseResponse{Message: message}
                    c.AbortWithStatusJSON(http.StatusInternalServerError, response)
                    return 
                }
                
                if !token.Valid {
                    message = "Internal error, Token not valid."
                    response := models.BaseResponse{Message: message}
                    c.AbortWithStatusJSON(http.StatusInternalServerError, response)
                    return 
                } 

                message = "Token success"   
                 
			} else { 
                message = "Bearer JWT Token not found."
                response := models.BaseResponse{Message: message}
                c.AbortWithStatusJSON(http.StatusInternalServerError, response)
                return 
            }
		} else { 
            message = "Token not found."
            response := models.BaseResponse{Message: message}
            c.AbortWithStatusJSON(http.StatusInternalServerError, response)
            return 
        }

        // Proceed to the next handler if authorized
        c.Next() 
    }
}