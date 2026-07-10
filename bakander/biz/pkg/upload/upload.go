package upload

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/disintegration/imaging"
	_ "github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type PublisherInterface interface {
	Publish(ctx context.Context, fileUploadTask FileUploadTask) error
}

type Service struct {
	publisher  PublisherInterface
	subscriber *Subscriber
	uploader   CloudUploader
}
type CloudUploader interface {
	Upload(ctx context.Context, localPath, remotePath, bucket string) error
}

func NewUploadService(publisher PublisherInterface, subscriber *Subscriber, cloudUploader CloudUploader) *Service {
	return &Service{
		publisher:  publisher,
		subscriber: subscriber,
		uploader:   cloudUploader,
	}
}

// 文档后缀
var docSuffixSet = map[string]struct{}{}

// 图片扩展名
var imageSuffixSet = map[string]struct{}{
	".jpg": {}, ".jpeg": {}, ".png": {},
	".gif": {}, ".bmp": {}, ".webp": {},
	".svg": {}, ".ico": {}, ".tiff": {},
	".tif": {},
}

// 视频扩展名
var videoSuffixSet = map[string]struct{}{
	".mp4": {}, ".avi": {}, ".wmv": {},
	".mkv": {}, ".webm": {}, ".m4v": {},
	".ogv": {}, ".ts": {}, ".m3u8": {},
	".w4v": {}, ".asf": {}, ".flv": {},
	".rmvb": {}, ".rm": {}, ".3gp": {},
	".vob": {}, ".wma": {}, ".mpeg": {},
	".mpg": {}, ".mov": {},
}

// 音频扩展名
var audioSuffixSet = map[string]struct{}{
	".mp3": {}, ".wav": {}, ".flac": {},
	".aac": {}, ".ogg": {}, ".m4a": {},
	".wma": {}, ".opus": {}, ".amr": {},
}

func (s *Service) UploadFile(ctx context.Context, fileHeader multipart.FileHeader) (string, string, error) {
	suffix, err := getFileSuffix(fileHeader.Filename)
	if err != nil {
		return "", "", err
	}
	// 声明节点
	sf, err := snowflake.NewNode(1)
	if err != nil {
		return "", "", err
	}
	taskId := sf.Generate().String()
	uploadPathBase := time.Now().Format("2006/01/02")
	task := FileUploadTask{
		FileTmpPath:     taskId + "." + suffix,
		CoverTmpPath:    taskId + "." + "png",
		FileUploadPath:  uploadPathBase + "/" + taskId + "." + suffix,
		CoverUploadPath: uploadPathBase + "/" + taskId + "." + "png",
	}
	uploadFile, err := os.Create(task.FileTmpPath)
	if err != nil {
		return "", "", err
	}
	defer uploadFile.Close()
	mpFile, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	defer mpFile.Close()
	_, err = uploadFile.ReadFrom(mpFile)
	if err != nil {
		return "", "", err
	}
	if err := s.publisher.Publish(ctx, task); err != nil {
		return "", "", err
	}
	urlPrefix := ""
	return urlPrefix + task.FileUploadPath, urlPrefix + task.CoverUploadPath, nil
}

func (s *Service) RunVideoUpload(ctx context.Context) error {
	return s.subscriber.Subscribe(ctx, "", func(body []byte) error {
		var task FileUploadTask
		if err := json.Unmarshal(body, &task); err != nil {
			return fmt.Errorf("cannot unmarshal json task: %v", err.Error())
		}

		if err := getVideoCover(task.FileTmpPath, task.CoverTmpPath); err != nil {
			return fmt.Errorf("get video cover err: fileTmpPath = %s", task.FileTmpPath)
		}
		var uploadErr error
		var wg sync.WaitGroup
		wg.Go(func() {
			defer func() {
				_ = os.Remove(task.CoverTmpPath)
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			if err := s.uploader.Upload(ctx, task.CoverTmpPath, task.CoverUploadPath, "video-thumb"); err != nil {
				log.Printf("cover upload err: %v", err)
				uploadErr = err
				return
			}
		})

		wg.Go(func() {
			defer func() {
				_ = os.Remove(task.FileTmpPath)
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			//suffix, err := getFileSuffix(task.FileTmpPath)
			//if err != nil {
			//	log.Printf("get video suffix err:videoTmpPath = %s", task.FileTmpPath)
			//	uploadErr = err
			//	return
			//}
			if err := s.uploader.Upload(ctx, task.FileTmpPath, task.FileUploadPath, "video"); err != nil {
				log.Printf("file upload err: %v", err)
				uploadErr = err
				return
			}
		})
		wg.Wait()

		if uploadErr != nil {
			return fmt.Errorf("nack err: %v", uploadErr)
		}

		return nil
	})

}
func (s *Service) RunImageUpload(ctx context.Context) error {
	return s.subscriber.Subscribe(ctx, "", func(body []byte) error {
		var task FileUploadTask
		if err := json.Unmarshal(body, &task); err != nil {
			return fmt.Errorf("cannot unmarshal json task: %v", err.Error())
		}
		defer func() {
			_ = os.Remove(task.FileTmpPath)
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		if err := s.uploader.Upload(ctx, task.FileTmpPath, task.FileUploadPath, "image"); err != nil {
			return fmt.Errorf("file upload err: %v", err)
		}

		return nil
	})
}
func (s *Service) RunDocUpload(ctx context.Context) error {
	return nil
}

// getFileSuffix
func getFileSuffix(fileName string) (suffix string, err error) {
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex < 0 {
		return "", fmt.Errorf("missing suffix : %v", lastDotIndex)
	}
	suffix = fileName[lastDotIndex+1:]
	suffix = strings.ToLower(suffix)
	return suffix, nil
}

// getVideoCover 保存视频封面图像到coverPath
func getVideoCover(videoPath, coverPath string) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{"gte(n,1)"}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).Run()
	if err != nil {
		return err
	}
	var img image.Image
	if img, err = imaging.Decode(buf); err != nil {
		return err
	}
	if err = imaging.Save(img, coverPath); err != nil {
		return err
	}
	return nil
}
