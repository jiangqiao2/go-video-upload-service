package repo

import (
	"context"
	"upload-service/ddd/domain/entity"
)

type TagRepository interface {
	ListAll(ctx context.Context) ([]*entity.TagEntity, error)
}
