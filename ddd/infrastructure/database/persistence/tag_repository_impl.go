package persistence

import (
	"context"
	"upload-service/ddd/domain/entity"
	"upload-service/ddd/domain/repo"
	"upload-service/ddd/infrastructure/database/dao"
)

type tagRepositoryImpl struct {
	dao *dao.TagDao
}

func NewTagRepository() repo.TagRepository {
	return &tagRepositoryImpl{dao: dao.NewTagDao()}
}

func (r *tagRepositoryImpl) ListAll(ctx context.Context) ([]*entity.TagEntity, error) {
	list, err := r.dao.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	entities := make([]*entity.TagEntity, 0, len(list))
	for _, t := range list {
		if t == nil {
			continue
		}
		entities = append(entities, entity.NewTagEntity(t.Id, t.TagUUID, t.Name, t.Code, t.Description))
	}
	return entities, nil
}
