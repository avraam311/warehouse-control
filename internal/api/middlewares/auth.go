package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wb-go/wbf/ginext"
)

var (
	ErrAuthHeaderMissing  = errors.New("authorization header missing or invalid")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidTokenClaims = errors.New("invalid token claims")
	ErrRoleNotFound       = errors.New("role not found in token")
	ErrAccessForbidden    = errors.New("access forbidden: insufficient role")
)

func abortWithError(c *ginext.Context, status int, err error) {
	c.AbortWithStatusJSON(status, ginext.H{"error": err.Error()})
}

func RoleBasedAuthMiddleware(jwtSecret string, allowedRoles []string) ginext.HandlerFunc {
	return func(c *ginext.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			abortWithError(c, http.StatusUnauthorized, ErrAuthHeaderMissing)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			abortWithError(c, http.StatusUnauthorized, ErrInvalidToken)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			abortWithError(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			abortWithError(c, http.StatusForbidden, ErrRoleNotFound)
			return
		}

		allowed := false
		for _, r := range allowedRoles {
			if r == role {
				allowed = true
				break
			}
		}
		if !allowed {
			abortWithError(c, http.StatusForbidden, ErrAccessForbidden)
			return
		}

		c.Set("role", role)
		c.Next()
	}
}
