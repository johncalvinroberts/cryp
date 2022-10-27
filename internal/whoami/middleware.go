// ref:
// https://gist.github.com/Goodnessuc/c92210cab062c541109e4fcf78bdfbe6
// https://blog.logrocket.com/jwt-authentication-go/
package whoami

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func handleBadToken(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
}

func VerifyWhoami(endpointHandler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		// no token in header
		if tokenHeader == "" {
			handleBadToken(c)
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)
			if !ok {
				return "", errors.New("failed to init token method")
			}
			return "", nil
		})
		// error parsing JWT
		if err != nil {
			handleBadToken(c)
		}
		// invalid token
		if !token.Valid {
			handleBadToken(c)
		}
		endpointHandler(c)
	}
}
