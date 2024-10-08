package utils

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

// ParseJsonArray 解析 JSON 数组到指定切片
func ParseJsonArray(pd datatypes.JSON, result interface{}) error {
	if err := json.Unmarshal(pd, result); err != nil {
		return fmt.Errorf("解析 JSON 时出错: %w", err)
	}
	return nil
}

// JsonArrayToStringSlice 将 JSON 数组转换为字符串切片
func JsonArrayToStringSlice(pd datatypes.JSON) ([]string, error) {
	var data [][]string
	if err := ParseJsonArray(pd, &data); err != nil {
		return nil, err
	}

	// 初始化 titles 切片
	titles := make([]string, len(data))
	for i, item := range data {
		if len(item) > 0 {
			titles[i] = item[0]
		}
	}

	// 使用切片的实际长度
	return titles[:len(data)], nil
}

// PromptJsonArrayToStringSlice 将 JSON 数组直接转换为字符串切片
func PromptJsonArrayToStringSlice(pd datatypes.JSON) ([]string, error) {
	var result []string
	if err := ParseJsonArray(pd, &result); err != nil {
		return nil, err
	}

	return result, nil
}
