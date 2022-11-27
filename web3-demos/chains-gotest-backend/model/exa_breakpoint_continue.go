package model

import (
	"chains-gotest-backend/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.BaseModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.BaseModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
