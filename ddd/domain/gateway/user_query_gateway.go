package gateway

import (
	"context"
	"upload-service/ddd/domain/vo"
)

type UserQueryGateway interface {
	GetUsersByUUIDs(ctx context.Context, userUUIDs []string) ([]*vo.UserSummary, error)
}
