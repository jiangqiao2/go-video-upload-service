package entity

type TagEntity struct {
	id          uint64
	tagUUID     string
	name        string
	code        string
	description string
}

func NewTagEntity(id uint64, tagUUID, name, code, description string) *TagEntity {
	return &TagEntity{
		id:          id,
		tagUUID:     tagUUID,
		name:        name,
		code:        code,
		description: description,
	}
}

func (t *TagEntity) Id() uint64 {
	return t.id
}

func (t *TagEntity) TagUUID() string {
	return t.tagUUID
}

func (t *TagEntity) Name() string {
	return t.name
}

func (t *TagEntity) Code() string {
	return t.code
}

func (t *TagEntity) Description() string {
	return t.description
}
