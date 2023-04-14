package service

import (
	"context"
	"duryun-blog/config"
	"duryun-blog/model"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

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

	key := "blog/" + fileHeader.Filename

	_, model.Err = client.Object.Put(context.Background(), key, file, nil)
	if model.Err != nil {
		return "", model.Err
	}

	return fmt.Sprintf("%s", client.Object.GetObjectURL(key)), nil
}
