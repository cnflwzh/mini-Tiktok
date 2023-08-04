package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mini-Tiktok/config"
)

const CoverParams = "?vframe/jpg/offset/1/w/360/h/480"

var FormUploader *storage.FormUploader
var mac = qbox.NewMac(config.KodoConfig.AccessKey, config.KodoConfig.SecretKey)

func init() {
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	FormUploader = storage.NewFormUploader(&cfg)
}

func VideoUploadToKodo(videoByte []byte, userId int64) (videoUrl string, coverUrl string, err error) {
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
