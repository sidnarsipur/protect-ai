package models

const (
	VideoFileTypeMp4  VideoFileType = "mp4"
	VideoFileTypeWebm VideoFileType = "webm"
	VideoFileTypeAvi  VideoFileType = "avi"
	VideoFileTypeMov  VideoFileType = "mov"
	VideoFileTypeMkv  VideoFileType = "mkv"
	VideoFileTypeFlv  VideoFileType = "flv"
	VideoFileTypeMpg  VideoFileType = "mpg"
	VideoFileTypeMpeg VideoFileType = "mpeg"
	VideoFileTypeM4v  VideoFileType = "m4v"
)

type VideoFileType string

type VideoFileTypes []VideoFileType
