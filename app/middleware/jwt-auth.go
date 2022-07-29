package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/helper"
	"github.com/tuananh3561/go_crm/app/service"
	"net/http"
	"os"
)

var (
	tokenToServer = os.Getenv("TOKEN_TO_SERVER")
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := getToken(c)
		if auth == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if auth == tokenToServer {

		} else {
			token, err := jwtService.ValidateToken(auth)
			if err == nil && token != nil && token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				var user = entity.User{
					Id:       int(claims["id"].(float64)),
					FullName: claims["fullname"].(string),
					Email:    claims["email"].(string),
					//RoleName: claims["role_name"].(map[string]interface{}),
				}
				c.Set("user", user)
				return
			} else {
				response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		}
	}
}

func getToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		authHeader = c.GetHeader("Token")
	}
	if authHeader == "" {
		if value, ok := c.GetQuery("Authorization"); ok {
			authHeader = value
		} else if value2, ok2 := c.GetQuery("Token"); ok2 {
			authHeader = value2
		} else if value3, ok3 := c.GetQuery("token"); ok3 {
			authHeader = value3
		}
	}
	return authHeader
}

func verifyToken() {

}
