package middleware

// const (
// 	authorizationHeaderKey  = "authorization"
// 	authorizationTypeBearer = "bearer"
// 	authorizationPayloadKey = "authorization_payload"
// )

// func AuthMiddleware(token token.Maker) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authorizationHeader := c.GetHeader(authorizationHeaderKey)

// 		if len(authorizationHeader) == 0 {
// 			err := errors.New("authorization header is not provided")
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
// 			return
// 		}

// 		fields := strings.Fields(authorizationHeader)
// 		if len(fields) < 2 {
// 			err := errors.New("invalid authorization header format")
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 			return
// 		}

// 		authorizationType := strings.ToLower(fields[0])
// 		if authorizationType != authorizationTypeBearer {
// 			err := fmt.Errorf("unsuported authorization type %s", authorizationType)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 			return
// 		}

// 		accessToken := fields[1]
// 		payload, err := token.VerifyToken(accessToken)

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 			return
// 		}
// 		c.Set(authorizationPayloadKey, payload)
// 		c.Next()
// 	}
// }
