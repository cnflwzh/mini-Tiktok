package utils

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"testing"
)

func TestVideoUploadToKodo(t *testing.T) {
	//读取视频文件为byte数组，传入函数，返回视频的url
	file, err := os.ReadFile("./biz/utils/test.mp4")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
		return
	}
	url, coverUrl, err := VideoUploadToKodo(file, 2234)
	if err != nil {
		t.Errorf("Error uploading video: %v", err)
		return
	}
	hlog.Info("url: ", url)
	hlog.Info("coverUrl: ", coverUrl)
}
