package vo

type UploadVideoInitCmd struct {
	FileName    string
	FileSize    int
	TotalChunks int
	UserUUID    string
	FileHash    string
}

type UploadChunkCmd struct {
	ChunkUUID       string
	UserUUID        string
	UploadVideoUUID string
	ChunkSize       int
	ChunkIndex      int
	ChunkData       []byte
	ChunkHash       string
}

type PresignChunkCmd struct {
	ChunkUUID       string
	UserUUID        string
	UploadVideoUUID string
	ChunkIndex      int
	ChunkSize       int
	ContentType     string
}

type CompleteChunkCmd struct {
	ChunkUUID       string
	UserUUID        string
	UploadVideoUUID string
	ChunkIndex      int
	ChunkSize       int
	ChunkHash       string
}

type MergeChunkCmd struct {
	UploadVideoUUID string
	UserUUID        string
}
