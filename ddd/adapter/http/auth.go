package http

import (
	"github.com/gin-gonic/gin"

	"upload-service/pkg/errno"
)

func ExtractUserUUID(ctx *gin.Context) (string, error) {
	userUUID := ctx.GetHeader("X-User-UUID")
	if userUUID == "" {
		return "", errno.ErrUnauthorized
	}
	return userUUID, nil
}
