package service

import (
	"fmt"
	"strings"
	"testing"
	"wechat_api/util"
)

func TestReply(t *testing.T) {
	fmt.Println(util.FilterEmptyString(strings.Split("圣诞 2", " ")))
	fmt.Println(strings.Split("圣诞 2", " "))

	content := "[圣诞节 2]"
	content = strings.Trim(content, "[")
	content = strings.Trim(content, "]")
	fmt.Println(content)
	fmt.Println(strings.Trim("[圣诞节]", "["))

}
