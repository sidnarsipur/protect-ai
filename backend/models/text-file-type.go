package models

const (
	TextFileTypeTxt     TextFileType = "txt"
	TextFileTypeWordOld TextFileType = "doc"
	TextFileTypeWord    TextFileType = "docx"
	TextFileTypePdf     TextFileType = "pdf"
	TextFileTypeOdt     TextFileType = "odt"
	TextFileTypeRtf     TextFileType = "rtf"
	TextFileTypeTex     TextFileType = "tex"
)

type TextFileType string

type TextFileTypes []TextFileType
