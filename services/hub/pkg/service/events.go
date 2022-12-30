package service

// UploadReady informs an client that an upload is ready to work with
type UploadReady struct {
	FileID    string
	Filename  string
	Timestamp string

	Message string
}
