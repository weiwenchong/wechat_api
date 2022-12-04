package christmas

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"wechat_api/cache"
	"wechat_api/log"
	"wechat_api/util"
)

const (
	christmas_id        = "christmas_%d"
	christmas_exec_lock = "christmas_exec_lock_%d"

	// redis hash set中基本信息
	owner = "owner_721"
	total = "total_721"
	done  = "done_721"
)

type Christmas struct {
	Id  int64
	Key string

	Total   int64
	Owner   string
	Done    int64
	Members map[string]string
}

func NewChristmas(Id int64) *Christmas {
	return &Christmas{Id: Id, Key: fmt.Sprintf(christmas_id, Id)}
}

func Start(name string, total int64) int64 {
	Id := genId()
	log.Infof("%d", Id)
	c := NewChristmas(Id)
	c.Total = total
	c.start(name)

	return Id
}

func (c *Christmas) start(name string) {
	cache.Client.HSet(c.Key, owner, name)
	cache.Client.HSet(c.Key, total, total)
	cache.Client.Expire(c.Key, 24*time.Hour)
	return
}

func (c *Christmas) AddMember(name string) error {
	cache.Client.HSet(c.Key, name, "")
	return nil
}

func (c *Christmas) LoadCurrentInfo() *Christmas {
	members := cache.Client.HGetAll(c.Key).Val()
	c.Owner = members[owner]
	c.Total, _ = strconv.ParseInt(members[total], 10, 64)
	c.Done, _ = strconv.ParseInt(members[done], 10, 64)
	delete(members, owner)
	delete(members, total)
	delete(members, done)
	c.Members = members
	return c
}

func (c *Christmas) GetSelfResult(name string) (res string, err error) {
	c.LoadCurrentInfo()
	if c.Done == 0 {
		err = c.executeAssignment()
		if err != nil {
			return
		}
	}

	res = c.Members[name]
	return
}

func (c *Christmas) GetCompleteResult() {

}

func (c *Christmas) executeAssignment() (err error) {
	// todo lock
	if int64(len(c.Members)) != c.Total {
		err = errors.New("人都没齐")
		return
	}
	if c.Done == 1 {
		return
	}

	members := make([]string, len(c.Members), len(c.Members))
	for k, _ := range c.Members {
		members = append(members, k)
	}
	for k, _ := range c.Members {
		if len(members) == 1 {
			c.Members[k] = members[0]
			break
		}
		ds := util.NewSliceDeleteParam(members, k)
		index := rand.Intn(len(ds))
		c.Members[k] = ds[index]

		members = util.NewSliceDeleteParam(members, ds[index])
	}

	for k, v := range c.Members {
		cache.Client.HSet(c.Key, k, v)
	}
	cache.Client.HSet(c.Key, done, "1")

	return
}

func genId() (id int64) {
	for {
		id = rand.Int63n(9000) + 1000
		if cache.Client.Exists(fmt.Sprintf(christmas_id, id)).Val() == 1 {
			break
		}
	}
	return id
}
