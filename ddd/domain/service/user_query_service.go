package service

import (
	"context"
	"upload-service/ddd/domain/entity"
	"upload-service/ddd/domain/gateway"
	"upload-service/ddd/domain/vo"
	"upload-service/pkg/logger"
)

type UserQueryService interface {
	GetUsersForVideos(ctx context.Context, videos []*entity.VideoEntity) (map[string]*vo.UserSummary, error)
	BatchQueryUsers(ctx context.Context, userUUIDs []string) (map[string]*vo.UserSummary, error)
}

type userQueryServiceImpl struct {
	gateway gateway.UserQueryGateway
}

func NewUserQueryService(gateway gateway.UserQueryGateway) UserQueryService {
	return &userQueryServiceImpl{
		gateway: gateway,
	}
}

func (s *userQueryServiceImpl) GetUsersForVideos(ctx context.Context, videos []*entity.VideoEntity) (map[string]*vo.UserSummary, error) {
	uuidSet := make(map[string]struct{}, len(videos))
	uuids := make([]string, 0, len(videos))
	for _, v := range videos {
		u := v.UserUUID()
		if u == "" {
			continue
		}
		if _, ok := uuidSet[u]; !ok {
			uuidSet[u] = struct{}{}
			uuids = append(uuids, u)
		}
	}
	return s.BatchQueryUsers(ctx, uuids)
}

func (s *userQueryServiceImpl) BatchQueryUsers(ctx context.Context, userUUIDs []string) (map[string]*vo.UserSummary, error) {
	res := make(map[string]*vo.UserSummary, len(userUUIDs))
	if len(userUUIDs) == 0 {
		return res, nil
	}
	infos, err := s.gateway.GetUsersByUUIDs(ctx, userUUIDs)
	if err != nil {
		logger.Warnf("GetUsersByUUIDs failed: %v", err)
		return res, nil
	}
	for _, info := range infos {
		if info == nil {
			continue
		}
		res[info.UserUUID] = &vo.UserSummary{
			UserUUID:  info.UserUUID,
			Account:   info.Account,
			AvatarUrl: info.AvatarUrl,
		}
	}
	return res, nil
}
