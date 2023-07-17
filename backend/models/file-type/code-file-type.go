package models

const (
	CodeFileTypeGo    CodeFileType = "go"
	CodeFileTypeJs    CodeFileType = "js"
	CodeFileTypeTs    CodeFileType = "ts"
	CodeFileTypePy    CodeFileType = "py"
	CodeFileTypeJava  CodeFileType = "java"
	CodeFileTypePhp   CodeFileType = "php"
	CodeFileTypeHtml  CodeFileType = "html"
	CodeFileTypeCss   CodeFileType = "css"
	CodeFileTypeC     CodeFileType = "c"
	CodeFileTypeSwift CodeFileType = "swift"
)

type CodeFileType string

type CodeFileTypes []CodeFileType
