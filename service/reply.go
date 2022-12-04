package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"wechat_api/log"
	"wechat_api/model"
	"wechat_api/service/christmas"
	"wechat_api/util"
)

var invalidParam = errors.New("错误的指令")

func ReplyMsg(receiveMsg *model.ReceiveMsg) (result string, err error) {
	fields := util.FilterEmptyString(strings.Split(receiveMsg.Content, " "))
	if len(fields) == 0 {
		err = invalidParam
		return
	}
	log.Infof("ReplyMsg receiveMsg fields %+v", fields)

	if fields[0] == "圣诞节" {
		var total int64
		if len(fields) != 2 {
			return "", invalidParam
		}
		total, err = strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return "", invalidParam
		}
		id := christmas.Start(receiveMsg.FromUserName, total)
		result = fmt.Sprintf("%d", id)
		return result, nil
	} else if fields[0] == "获取结果" {
		var id int64
		if len(fields) != 2 {
			return "", invalidParam
		}
		id, err = strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return "", invalidParam
		}
		result, err = christmas.NewChristmas(id).GetSelfResult(receiveMsg.FromUserName)
		if err != nil {
			return "", err
		}
		return
	} else if id, err := strconv.ParseInt(fields[0], 10, 64); err == nil {
		c := christmas.NewChristmas(id)
		err = c.AddMember(receiveMsg.FromUserName)
		if err != nil {
			return "", err
		}
		c.LoadCurrentInfo()
		for k, _ := range c.Members {
			result = result + k + "\n"
		}
		return result, nil
	}

	return
}
