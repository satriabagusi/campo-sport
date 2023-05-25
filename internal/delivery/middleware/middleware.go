package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" && (c.Request.URL.Path == "/auth/login/" || c.Request.URL.Path == "/api/v1/users/register") {
			c.Next()
			return
		}

		//Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		//Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		//Parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			//Provide your own secret key or verification logic here
			//You can retrieve the secret key from a config file environtment variable
			//For example, you can use the same key to sign and verivy the token
			secretKey := utility.GetEnv("SECRET_KEY")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		//check if the token is valid
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		//Set the user information iun the context for further processing
		//For example, you can extract the user ID from the token and set it in the context
		//The user information can be retrieved in the handler functions using c.Getstring("userID")
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		userIdClaims, ok := claims["userID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid userId claim"})
			c.Abort()
			return
		}

		userID := fmt.Sprintf("%.0f", userIdClaims) //Convert float64 to string
		c.Set("userID", userID)

		//continue to the next middleware or handler
		c.Next()
	}
}
