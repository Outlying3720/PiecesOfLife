package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"demo/models"
)

func ExtractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func ValidateToken(c *gin.Context) (string, error) {
	tk := ExtractToken(c)
	if tk == "" {
		return "", fmt.Errorf("header without token")
	}

	user, err := models.ExamToken(tk)

	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		fmt.Println("welcome user:", user)
		ctx.Set("auth", user)
		ctx.Next()
	}
}
