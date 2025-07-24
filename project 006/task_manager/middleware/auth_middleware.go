package middleware

import (
	"fmt" // For formatted errors
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct{} 

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"` 
	jwt.RegisteredClaims
}

func (a *AuthMiddleware) GenerateJWT(username string, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role, 
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "task_manager",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Authenticate middleware: validates token and stores claims in context
func (a *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			// log.Printf("Token validation error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}


		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next() 
	}
}

func (a *AuthMiddleware) Authorize(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the role from the context, which was set by the Authenticate middleware
		role, exists := c.Get("role")
		if !exists {
			// This means Authenticate middleware didn't run or failed to set the role
			// This should theoretically not happen if Authenticate is applied first
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User role not found in context"})
			return
		}

		userRole, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type in context"})
			return
		}

		// Check if the user's role is among the required roles
		isAuthorized := false
		for _, requiredRole := range requiredRoles {
			if userRole == requiredRole {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient permissions"})
			return
		}

		c.Next() 
	}
}