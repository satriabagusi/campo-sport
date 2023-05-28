package middleware

// func Authentication2() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		verifyToken, err := token.VerifyToken(c)
// 		_ = verifyToken

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthenticated",
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		c.Set("userID", verifyToken)
// 		c.Next()
// 	}
// }
