package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	AccessKey = utils.AccessKey
	SecretKey = utils.SecretKey
	Bucket    = utils.Bucket
	ImgUrl    = utils.QiniuServer
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	region, _ := storage.GetRegion(AccessKey, Bucket)
	cfg := storage.Config{
		Zone:          region,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCSE
}
