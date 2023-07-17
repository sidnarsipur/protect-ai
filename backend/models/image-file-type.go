package models

const (
	ImageFileTypePng  ImageFileType = "png"
	ImageFileTypeJpg  ImageFileType = "jpg"
	ImageFileTypeJpeg ImageFileType = "jpeg"
	ImageFileTypeGif  ImageFileType = "gif"
	ImageFileTypeSvg  ImageFileType = "svg"
	ImageFileTypeWebp ImageFileType = "webp"
	ImageFileTypeBmp  ImageFileType = "bmp"
	ImageFileTypeIco  ImageFileType = "ico"
	ImageFileTypeTiff ImageFileType = "tiff"
	ImageFileTypeTif  ImageFileType = "tif"
	ImageFileTypeAvif ImageFileType = "avif"
	ImageFileTypeAi   ImageFileType = "ai"
)

type ImageFileType string

type ImageFileTypes []ImageFileType
