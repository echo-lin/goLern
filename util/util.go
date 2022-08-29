package util

import (
	"golern/dto"
	"math/rand"
)

func RandomString(n int) string {
	var letters = []byte("ASDFGHJKlQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// UnlimitedClass
// @Description 生成无限级分类树
// @Date 2022-08-28 20:29:53
// @param data
// @return map[int]*dto.ChapterDto
func UnlimitedClass(data map[int]*dto.ChapterDto) map[int]*dto.ChapterDto {
	unlClass := make(map[int]*dto.ChapterDto)
	for key, cur := range data {
		if value, ok := data[cur.Pid]; ok {
			// 找到每个节点的子节点
			value.Children = append(value.Children, cur)
			// 更新原数据
			data[cur.Pid] = value
		} else {
			// 只获取顶层节点
			unlClass[cur.ID] = data[key]
		}
	}
	return unlClass
}
