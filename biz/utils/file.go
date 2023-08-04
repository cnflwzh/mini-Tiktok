package utils

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

// ReadFile 将文件读取为byte数组
func ReadFile(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			hlog.Error("close file error", err)
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GenerateFileName 生成上传到七牛云的文件名
func GenerateFileName(userId int64, extension string) string {
	var str strings.Builder
	t := time.Now()
	fileUUID := uuid.New().String()

	fmt.Fprintf(&str, "/%s/%02d/%02d/%d-%d-%s.%s",
		t.Format("2006"), t.Month(), t.Day(), t.Unix(), userId, fileUUID, extension)
	return str.String()
}

// 判断是否为视频文件
func IsMp4Video(b []byte) bool {
	kind, err := filetype.Match(b)
	if err != nil {
		return false
	}
	return kind.MIME.Value == "video/mp4"
}
