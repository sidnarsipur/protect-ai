package models

const (
	AudioFileTypeMp3  AudioFileType = "mp3"
	AudioFileTypeWav  AudioFileType = "wav"
	AudioFileTypeFlac AudioFileType = "flac"
	AudioFileTypeOgg  AudioFileType = "ogg"
	AudioFileTypeAiff AudioFileType = "aiff"
	AudioFileTypeAif  AudioFileType = "aif"
	AudioFileTypeM4a  AudioFileType = "m4a"
	AudioFileTypeWma  AudioFileType = "wma"
	AudioFileTypeAac  AudioFileType = "aac"
)

type AudioFileType string

type AudioFileTypes []AudioFileType
