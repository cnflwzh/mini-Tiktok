package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mini-Tiktok/config"
	"strings"
	"time"
)

const CoverParams = "?vframe/jpg/offset/1/w//h/480"

var FormUploader *storage.FormUploader
var mac = qbox.NewMac(config.KodoConfig.AccessKey, config.KodoConfig.SecretKey)

func init() {

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	FormUploader = storage.NewFormUploader(&cfg)
}

func GenerateFileName(userId int64, extension string) string {
	var str strings.Builder
	t := time.Now()
	fileUUID := uuid.New().String()

	fmt.Fprintf(&str, "/%s/%02d/%02d/%d-%d-%s.%s",
		t.Format("2006"), t.Month(), t.Day(), t.Unix(), userId, fileUUID, extension)
	return str.String()
}

func isVideo(b []byte) bool {
	kind, err := filetype.Match(b)
	if err != nil {
		return false
	}
	return kind.MIME.Value == "video/mp4"
}

func VideoUploadToKodo(videoByte []byte, userId int64) (videoUrl string, coverUrl string, err error) {
	videoFormat := isVideo(videoByte)
	if videoFormat == false {
		return "", "", fmt.Errorf("上传的不是视频或格式暂不支持")
	}
	key := GenerateFileName(userId, "mp4") // 使用生成的文件名
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", config.KodoConfig.Bucket, key),
	}
	upToken := putPolicy.UploadToken(mac)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(videoByte))
	dataReader := bytes.NewReader(videoByte)
	err = FormUploader.Put(context.Background(), &ret, upToken, key, dataReader, dataLen, &putExtra)
	if err != nil {
		return "", "", fmt.Errorf("error uploading video %s: %w", key, err)
	}

	videoUrl = fmt.Sprintf("%s%s", config.KodoConfig.Domain, ret.Key)
	coverUrl = fmt.Sprintf("%s%s%s", config.KodoConfig.Domain, ret.Key, CoverParams)

	return videoUrl, coverUrl, nil
}
