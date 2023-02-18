package middleware

import (
	"errors"
	"go-template/model"
	"go-template/sdk/apires"
	sdk_jwt "go-template/sdk/jwt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			apires.FailOrError(c, 
				http.StatusUnauthorized, "Not authorized, login/register first",
				errors.New("\"Authorization\" header is not set"),
			)
			c.Abort()
			return		
		}

		// helps shorten code
		errResp := func(code int, msg string, err string){
			apires.FailOrError(c, http.StatusUnauthorized, msg,errors.New(err))
			c.Abort()
		}
		msgUnauthorized := "You are not authorized to access this resource"

		//splits "<type> <token>"
		parts := strings.Split(authorization, " ")
		if len(parts) < 2 {
			errResp(http.StatusUnauthorized, msgUnauthorized, "invalid value")
			return
		}

		jwtKey := os.Getenv("JWT_KEY")
		var claims model.UserClaims

		switch parts[0]{
			case "Bearer" :
				if err := sdk_jwt.DecodeToken(parts[1], &claims, jwtKey); err != nil {
					errResp(http.StatusUnauthorized, msgUnauthorized, err.Error())
					return
				}
				c.Set("user", claims)

			default :
				errResp(http.StatusUnauthorized, msgUnauthorized, "invalid token type")
				return
		}
		c.Next()
	}
}