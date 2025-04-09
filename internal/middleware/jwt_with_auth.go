package middleware

import (
	"net/http"
	"strings"

	"guilt-type-service/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(authClient auth.AuthClient, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Локальная валидация токена
		if err := ValidateJWT(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Проверка через внешний Java Auth сервис
		userInfo, err := authClient.VerifyUser(token)
		if err != nil || !userInfo.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "External auth failed"})
			return
		}

		if userInfo.Role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		// Можно сохранить userInfo в context
		c.Set("user", userInfo)
		c.Next()
	}
}
