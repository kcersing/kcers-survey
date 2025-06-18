package utils

import (
	"testing"
)

func TestCreateCn(t *testing.T) {

	// 测试函数 CreateCn 的结果是否与预期一致
	result := CreateCn()

	t.Errorf("CreateCn() = %s;length : %v", result, len(result))

}
