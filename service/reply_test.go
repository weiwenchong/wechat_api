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

}
