package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mini-Tiktok/config"
	"time"
)

func GenerateFileName(userId int64, extension string) string {
	t := time.Now()
	fileUUID := uuid.New().String()

	fileName := fmt.Sprintf("%s/%02d/%02d/%d-%d-%s.%s",
		t.Format("2006"), t.Month(), t.Day(), t.Unix(), userId, fileUUID, extension)
	return fileName
}

func VideoUploadToKodo(videoByte []byte, userId int64) (videoUrl string, coverUrl string, err error) {
	mac := qbox.NewMac(config.KodoConfig.AccessKey, config.KodoConfig.SecretKey)
	key := GenerateFileName(userId, "mp4") // 使用生成的文件名
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", config.KodoConfig.Bucket, key),
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(videoByte))
	dataReader := bytes.NewReader(videoByte)
	err = formUploader.Put(context.Background(), &ret, upToken, key, dataReader, dataLen, &putExtra)
	if err != nil {
		return "", "", err
	}

	videoUrl = fmt.Sprintf("http://ryi1c8rs5.hn-bkt.clouddn.com/%s", ret.Key)
	coverUrl = fmt.Sprintf("http://ryi1c8rs5.hn-bkt.clouddn.com/%s", ret.Key, "?vframe/jpg/offset/1/w/360/h/480")

	return videoUrl, coverUrl, nil
}
