package util

import "strings"

func NewSliceDeleteParam(data []string, delete string) (res []string) {
	for _, v := range data {
		if v != delete {
			res = append(res, v)
		}
	}
	return
}

func FilterEmptyString(input []string) []string {
	ret := make([]string, 0, len(input))
	for _, item := range input {
		cur := strings.TrimSpace(item)
		if len(cur) == 0 {
			continue
		}
		ret = append(ret, item)
	}
	return ret
}
