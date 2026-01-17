package vo

type PublishVideoCmd struct {
	UploadVideoUUID  string
	Title            string
	Description      string
	Tags             []string
	CoverURL         string
	UserUUID         string
	TargetResolution string
	TargetBitrate    string
}
