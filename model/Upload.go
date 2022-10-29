package model

import (
	"context"
	"duryun-blog/config"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

//var Zone = utils.Zone
//var AccessKey = utils.AccessKey
//var SecretKey = utils.SecretKey
//var Bucket = utils.Bucket
//var ImgUrl = utils.QiniuSever

var (
	BucketURL = config.BucketURL
	SecretID  = config.SecretID
	SecretKey = config.SecretKey
)

func UpLoadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	u, _ := url.Parse(BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{

			SecretID:  os.Getenv(SecretID),
			SecretKey: os.Getenv(SecretKey),
		},
	})

	key := "test/" + fileHeader.Filename

	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", client.Object.GetObjectURL(key)), nil
}

//func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
//	putPolicy := storage.PutPolicy{
//		Scope: Bucket,
//	}
//	mac := qbox.NewMac(AccessKey, SecretKey)
//	upToken := putPolicy.UploadToken(mac)
//
//	cfg := storage.Config{
//		Zone:          selectZone(Zone),
//		UseCdnDomains: false,
//		UseHTTPS:      false,
//	}
//
//	putExtra := storage.PutExtra{}
//
//	formUploader := storage.NewFormUploader(&cfg)
//	ret := storage.PutRet{}
//
//	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	url := ImgUrl + ret.Key
//	return url, errmsg.SUCCSE
//
//}
//
//
//
//func selectZone(id int) *storage.Zone {
//	switch id {
//	case 1:
//		return &storage.ZoneHuadong
//	case 2:
//		return &storage.ZoneHuabei
//	case 3:
//		return &storage.ZoneHuanan
//	default:
//		return &storage.ZoneHuadong
//	}
//}
