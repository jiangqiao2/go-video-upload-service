package grpc

import (
	"context"
	"upload-service/ddd/domain/gateway"
	"upload-service/ddd/domain/vo"
)

type userQueryGatewayImpl struct {
	client *UserServiceClient
}

func NewUserQueryGateway(client *UserServiceClient) gateway.UserQueryGateway {
	return &userQueryGatewayImpl{client: client}
}

func (g *userQueryGatewayImpl) GetUsersByUUIDs(ctx context.Context, userUUIDs []string) ([]*vo.UserSummary, error) {
	if g.client == nil {
		return nil, nil
	}
	infos, err := g.client.GetUsersByUUIDs(ctx, userUUIDs)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	res := make([]*vo.UserSummary, 0, len(infos))
	for _, info := range infos {
		if info == nil {
			continue
		}
		res = append(res, &vo.UserSummary{
			UserUUID:  info.UserUuid,
			Account:   info.Account,
			AvatarUrl: info.AvatarUrl,
		})
	}
	return res, nil
}
