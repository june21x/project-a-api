package middleware

// import (
// 	"log"
// 	"time"

// 	jwt "github.com/appleboy/gin-jwt/v2"
// 	"github.com/gin-gonic/gin"
// 	playerservice "github.com/june21xproject-a-api/internal_service/player"
// )

// type login struct {
// 	Playername string `form:"playerName" json:"playerName" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// var identityKey = "id"

// func AuthMiddleware() gin.HandlerFunc {
// 	// the jwt middleware
// 	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
// 		Realm:       "test zone",
// 		Key:         []byte("secret key"),
// 		Timeout:     time.Hour,
// 		MaxRefresh:  time.Hour,
// 		IdentityKey: identityKey,
// 		PayloadFunc: func(data interface{}) jwt.MapClaims {
// 			if v, ok := data.(*playerservice.Player); ok {
// 				return jwt.MapClaims{
// 					identityKey: v.PlayerName,
// 				}
// 			}
// 			return jwt.MapClaims{}
// 		},
// 		IdentityHandler: func(c *gin.Context) interface{} {
// 			claims := jwt.ExtractClaims(c)
// 			return &playerservice.Player{
// 				PlayerName: claims[identityKey].(string),
// 			}
// 		},
// 		Authenticator: func(c *gin.Context) (interface{}, error) {
// 			var loginVals login
// 			if err := c.ShouldBind(&loginVals); err != nil {
// 				return "", jwt.ErrMissingLoginValues
// 			}
// 			playerID := loginVals.Playername
// 			password := loginVals.Password

// 			if (playerID == "admin" && password == "admin") || (playerID == "test" && password == "test") {
// 				return &playerservice.Player{
// 					PlayerName:  playerID,
// 					LastName:  "Bo-Yi",
// 					FirstName: "Wu",
// 				}, nil
// 			}

// 			return nil, jwt.ErrFailedAuthentication
// 		},
// 		Authorizator: func(data interface{}, c *gin.Context) bool {
// 			if v, ok := data.(*playerservice.Player); ok && v.PlayerName == "admin" {
// 				return true
// 			}

// 			return false
// 		},
// 		Unauthorized: func(c *gin.Context, code int, message string) {
// 			c.JSON(code, gin.H{
// 				"code":    code,
// 				"message": message,
// 			})
// 		},
// 		// TokenLookup is a string in the form of "<source>:<name>" that is used
// 		// to extract token from the request.
// 		// Optional. Default value "header:Authorization".
// 		// Possible values:
// 		// - "header:<name>"
// 		// - "query:<name>"
// 		// - "cookie:<name>"
// 		// - "param:<name>"
// 		TokenLookup: "header: Authorization, query: token, cookie: jwt",
// 		// TokenLookup: "query:token",
// 		// TokenLookup: "cookie:token",

// 		// TokenHeadName is a string in the header. Default value is "Bearer"
// 		TokenHeadName: "Bearer",

// 		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
// 		TimeFunc: time.Now,
// 	})

// 	if err != nil {
// 		log.Fatal("JWT Error:" + err.Error())
// 	}

// 	return authMiddleware.MiddlewareFunc()

// }
