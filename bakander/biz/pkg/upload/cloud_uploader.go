package upload

import (
	"context"
	"path"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type OSSCloudUploader struct {
	//client *oss.Client
}

func (u *OSSCloudUploader) Upload(ctx context.Context, localPath, remotePath, bucket string) error {
	return nil
}

type TOSCloudUploader struct {
	//client *tos.Client
}

func (u *TOSCloudUploader) Upload(ctx context.Context, localPath, remotePath, bucket string) error {
	return nil
}

type FastdfsUploader struct {
}

func (u *FastdfsUploader) Upload(ctx context.Context, localPath, remotePath, bucket string) error {

	cli, err := client.NewClient()
	if err != nil {
		return err
	}

	req := protocol.AcquireRequest()
	rsp := protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(rsp)
	}()

	//req.SetFormDataFromValues(url.Values{
	//	"path":   []string{remotePath},
	//	"scene":  []string{bucket},
	//	"output": []string{"json"},
	//})
	//req.SetFormData(map[string]string{
	//	"path":   path.Dir(remotePath),
	//	"scene":  bucket,
	//	"output": "json",
	//})

	//args := req.PostArgs()
	//args.Add("path", remotePath)
	//args.Add("scene", bucket)
	//args.Add("output", "json")

	req.SetFile("file", localPath)

	req.SetMultipartFormData(map[string]string{
		"path":   path.Dir(remotePath),
		"scene":  bucket,
		"output": "json",
	})
	hlog.Info(remotePath)
	hlog.Info(bucket)

	req.SetRequestURI("http://localhost:9109/upload")
	req.SetMethod(consts.MethodPost)

	err = cli.Do(ctx, req, rsp)
	if err != nil {
		return err
	}
	if rsp.StatusCode() == 200 {
		var frsp fastdfsRsp
		err := json.Unmarshal(rsp.Body(), &frsp)
		if err != nil {
			return err
		}
		hlog.Infof("err : %v ,Body : %v", err, string(rsp.Body()))
	}

	return nil
}

type fastdfsRsp struct {
	Domain  string `json:"domain"`
	Md5     string `json:"md5"`
	Mtime   int    `json:"mtime"`
	Path    string `json:"path"`
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Scene   string `json:"scene"`
	Scenes  string `json:"scenes"`
	Size    int    `json:"size"`
	Src     string `json:"src"`
	Url     string `json:"url"`
}
