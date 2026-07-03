package survey

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// 需要校验为数字的类型
var numericTypes = map[string]bool{
	"number": true, "slider": true, "rate": true, "nps": true,
}

// validateAnswer 按题目类型校验答案格式
func validateAnswer(qType string, value []string) error {
	if len(value) == 0 {
		return errors.New("答案不能为空")
	}

	// 数字类：value[0] 必须是合法数字
	if numericTypes[qType] {
		if _, err := strconv.ParseFloat(value[0], 64); err != nil {
			return errors.Errorf("答案格式错误: %s 应为数字, 实际为 %s", qType, value[0])
		}
		return nil
	}

	// 矩阵题：value[0] 必须是合法 JSON
	if qType == "matrix" {
		var tmp map[string]interface{}
		if err := json.Unmarshal([]byte(value[0]), &tmp); err != nil {
			return errors.Errorf("矩阵答案格式错误: %s", value[0])
		}
		return nil
	}

	// 日期：YYYY-MM-DD
	if qType == "date" {
		if ok, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, value[0]); !ok {
			return errors.Errorf("日期格式错误: %s (应为 YYYY-MM-DD)", value[0])
		}
		return nil
	}

	// 日期时间：YYYY-MM-DD HH:mm:ss
	if qType == "datetime" {
		if ok, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`, value[0]); !ok {
			return errors.Errorf("日期时间格式错误: %s (应为 YYYY-MM-DD HH:mm:ss)", value[0])
		}
		return nil
	}

	// 排序题：必须含 > 分隔符
	if qType == "ranking" {
		if !strings.Contains(value[0], ">") {
			return errors.Errorf("排序答案格式错误: %s (应包含 > 分隔符)", value[0])
		}
		return nil
	}

	return nil
}

// shouldStoreAsText 判断答案是否应存入 answer_text 字段
func shouldStoreAsText(reqType string) bool {
	return reqType == "input" || reqType == "text"
}
