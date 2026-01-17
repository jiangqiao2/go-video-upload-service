package vo

type MergeChunkResult struct {
	Status            string
	UploadVideoUUID   string
	ShouldEnqueueTask bool
}
