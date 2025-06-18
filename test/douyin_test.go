package test

import (
	"github.com/kcersing/dy"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestToken(t *testing.T) {

	var client dy.DyInterface = &dy.Dy{
		Config: dy.Config{
			AppId:     "aw3v00he8vmgvzty",
			AppSecret: "a853336c7feaa588a62f0250c465b2a1",
			AccountId: "7218069838395082791",
		},
	}

	dto := &dy.GenAuthWithBindValidUrlDto{
		Timestamp:      strconv.FormatInt(time.Now().Unix(), 10),
		Extra:          "",
		SolutionKey:    "4",
		OutShopId:      "aw3v00he8vmgvzty",
		PermissionKeys: []string{"1", "2", "5", "6", "9", "10", "12", "15", "16"},
	}

	result, _ := client.GenAuthValidUrl(dto)

	log.Printf("result: %v", result)
}

//func TestGetAccessTokenTimeout(t *testing.T) {
//	// 创建模拟客户端
//	mockClient := new(MockClient)
//
//	// 模拟客户端超时错误
//	mockClient.On("Do", context.Background(), mock.Anything, mock.Anything).Return(context.DeadlineExceeded)
//
//	// 调用目标函数
//	_, err := GetAccessToken()
//
//	// 断言错误为超时错误
//	assert.Equal(t, context.DeadlineExceeded, err)
//
//	// 验证模拟客户端的调用
//	mockClient.AssertExpectations(t)
//}
