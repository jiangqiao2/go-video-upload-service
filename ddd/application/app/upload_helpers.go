package app

import (
	"fmt"
	"time"
	"upload-service/ddd/application/dto"
	"upload-service/ddd/domain/entity"
	"upload-service/ddd/domain/vo"
)

// attachPresignForChunks maps entity presign info onto response DTOs.
func attachPresignForChunks(uploadVideoEntity *entity.UploadVideoEntity, res *dto.UploadVideoDto, chunks []*entity.UploadChunkEntity) {
	if uploadVideoEntity == nil || res == nil || len(res.UploadChunks) == 0 {
		return
	}
	entities := make(map[string]*entity.UploadChunkEntity, len(chunks))
	for _, c := range chunks {
		if c == nil {
			continue
		}
		entities[c.ChunkUUID()] = c
	}
	for i := range res.UploadChunks {
		ch := &res.UploadChunks[i]
		ent := entities[ch.ChunkUUID]
		key := ""
		if ent != nil {
			key = ent.StoragePath()
		}
		if key == "" {
			key = fmt.Sprintf("%s%d", uploadVideoEntity.ChunkStoragePath(), ch.ChunkIndex)
		}
		ch.StoragePath = key
		if ch.Status == vo.UploadChunkStatusCompleted.Value() || key == "" {
			continue
		}
		if ent != nil {
			ch.PutURL = ent.PutURL()
			if exp := ent.PresignExpiredAt(); exp != nil {
				remaining := int(exp.Sub(time.Now()).Seconds())
				if remaining < 0 {
					remaining = 0
				}
				ch.ExpiresInSec = remaining
			}
		}
	}
}
