package app

import (
	"context"
	"sync"
	"upload-service/ddd/application/cqe"
	"upload-service/ddd/application/dto"
	"upload-service/ddd/domain/repo"
	"upload-service/ddd/infrastructure/database/persistence"
)

type TagApp interface {
	ListTags(ctx context.Context, req *cqe.ListTagsReq) (*dto.TagListDto, error)
}

type tagAppImpl struct {
	tagRepo repo.TagRepository
}

var (
	defaultTagApp TagApp
	tagAppOnce    sync.Once
)

func DefaultTagApp() TagApp {
	tagAppOnce.Do(func() {
		defaultTagApp = &tagAppImpl{
			tagRepo: persistence.NewTagRepository(),
		}
	})
	return defaultTagApp
}

func (a *tagAppImpl) ListTags(ctx context.Context, req *cqe.ListTagsReq) (*dto.TagListDto, error) {
	list, err := a.tagRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	dtos := make([]*dto.TagDto, 0, len(list))
	for _, t := range list {
		dtos = append(dtos, &dto.TagDto{Id: t.Id(), TagUUID: t.TagUUID(), Name: t.Name(), Code: t.Code(), Description: t.Description()})
	}
	return dto.NewTagListDto(dtos), nil
}
