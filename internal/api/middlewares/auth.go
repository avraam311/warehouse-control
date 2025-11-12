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
	ErrUserIDNotFound     = errors.New("user_id not found in token")
	ErrRoleNotFound       = errors.New("role not found in token")
	ErrAccessForbidden    = errors.New("access forbidden: insufficient permission")
)

var rolePermissions = map[string]map[string]struct{}{
	"admin": {
		"POST:/warehouse-control/api/auth/register": {},
		"POST:/warehouse-control/api/auth/login/":   {},
		"POST:/warehouse-control/api/items/":        {},
		"GET:/warehouse-control/api/items/":         {},
		"PUT:/warehouse-control/api/items/:id":      {},
		"DELETE:/warehouse-control/api/items/:id":   {},
	},
	"manager": {
		"POST:/warehouse-control/api/auth/login/": {},
		"GET:/warehouse-control/api/items/":       {},
		"PUT:/warehouse-control/api/items/:id":    {},
	},
	"viewer": {
		"GET:/warehouse-control/api/items/": {},
	},
}

func abortWithError(c *ginext.Context, status int, err error) {
	c.AbortWithStatusJSON(status, ginext.H{"error": err.Error()})
}

func RoleBasedAuthMiddleware(jwtSecret string) ginext.HandlerFunc {
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

		role, ok := claims["Role"].(string)
		if !ok {
			abortWithError(c, http.StatusForbidden, ErrRoleNotFound)
			return
		}

		method := c.Request.Method
		path := c.FullPath()

		key := method + ":" + path
		perms, ok := rolePermissions[role]
		if !ok {
			abortWithError(c, http.StatusForbidden, ErrAccessForbidden)
			return
		}
		if _, allowed := perms[key]; !allowed {
			abortWithError(c, http.StatusForbidden, ErrAccessForbidden)
			return
		}

		userIDFloat, ok := claims["UserID"].(float64)
		if !ok {
			abortWithError(c, http.StatusForbidden, ErrUserIDNotFound)
			return
		}
		userID := uint(userIDFloat)

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Next()
	}
}
