package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"kcers-survey/biz/dal/config"
	"log"
	"mime/multipart"
	"net/url"

	"strings"
	"time"
)

var (
	Client *minio.Client
	err    error
)

func InitMinio() {
	ctx := context.Background()
	Client, err = minio.New(config.GlobalServerConfig.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.GlobalServerConfig.Minio.AccessKeyID, config.GlobalServerConfig.Minio.SecretAccessKey, ""),
		Secure: config.GlobalServerConfig.Minio.UseSSL,
	})

	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}

	MakeBucket(ctx, config.GlobalServerConfig.Minio.VideoBucketName)
	MakeBucket(ctx, config.GlobalServerConfig.Minio.ImgBucketName)
}

// MakeBucket create a bucket with a specified name
func MakeBucket(ctx context.Context, bucketName string) {
	hlog.Info(bucketName)
	exists, err := Client.BucketExists(ctx, bucketName)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("已成功创建 %v\n", bucketName)
	}
}

// PutToBucket put the file into the bucket by *multipart.FileHeader
func PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	info, err = Client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	fileObj.Close()
	return info, err
}

// GetObjURL get the original link of the file in minio
func GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	u, err = Client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}

// PutToBucketByBuf put the file into the bucket by *bytes.Buffer
func PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByFilePath put the file into the bucket by filepath
func PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = Client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}

// NewFileName Splicing user_id and time to make unique filename
func NewFileName(user_id, time int64) string {
	if user_id == 0 {
		return fmt.Sprintf("%d", time)
	}
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = string(c.URI().Scheme())
	u.Host = string(c.URI().Host())
	//u.Host = "127.0.0.1:9039"
	u.Path = "/src" + u.Path
	return u.String()
}
