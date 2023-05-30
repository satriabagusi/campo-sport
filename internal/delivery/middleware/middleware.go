package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/pkg/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")

		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		user, err := token.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		c.Set("userinfo", user)

		c.Next()
	}
}

// func Authentication() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		//Get the Authorization header
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
// 			c.Abort()
// 			return
// 		}

// 		//Extract the token from the header
// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		if tokenString == authHeader {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 			c.Abort()
// 			return
// 		}

// 		//Parse the token
// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			secretKey := utility.GetEnv("SECRET_KEY")
// 			return []byte(secretKey), nil
// 		})

// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		//check if the token is valid
// 		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}
// 		claims, ok := token.Claims.(jwt.MapClaims)

// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
// 			c.Abort()
// 			return
// 		}

// 		userIdClaims, ok := claims["userID"].(float64)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid userId claim"})
// 			c.Abort()
// 			return
// 		}

// 		userID := fmt.Sprintf("%.0f", userIdClaims)
// 		c.Set("userID", userID)

// 		c.Next()
// 	}
// }
