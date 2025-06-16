package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

func doClient() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	// Do
	req, res := &protocol.Request{}, &protocol.Response{}
	req.SetOptions(config.WithRequestTimeout(5 * time.Second))
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://localhost:8888/get")
	err = c.Do(context.Background(), req, res)
	fmt.Printf("Status Code = %v,resp = %v,err = %+v", res.StatusCode(), string(res.Body()), err)

	if err != nil {
		return
	}

	// close idle connections
	c.CloseIdleConnections()
}

// 单一上传
func singleFile() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/singleFile")
	req.SetFile("file", "your file path")

	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Println(err, string(res.Body()))
}
