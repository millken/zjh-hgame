package rd

import (
	"fmt"
	"log"

	"github.com/millken/zjh-hgame/common"

	"gopkg.in/redis.v3"
)

var (
	rc  *redis.Client
	err error
)

func Boot(addr string) {
	rc = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err = rc.Ping().Result()
	if err != nil {
		log.Fatalf("connect redis server error: %s", err)
	}
}

func Client() *redis.Client {
	return rc
}

func SetUserToken(uid int, ghostId string, hallToken string) error {
	if err = rc.Set(fmt.Sprintf("u:%d:ghostId", uid), ghostId, 0).Err(); err != nil {
		return fmt.Errorf("set ghostId err %s", err)
	}
	if err = rc.Set(fmt.Sprintf("hallToken:%s:uid", hallToken), uid, 0).Err(); err != nil {
		return fmt.Errorf("set hallTokenuid err %s", err)
	}
	if err = rc.Set(fmt.Sprintf("u:%d:hallToken", uid), hallToken, 0).Err(); err != nil {
		return fmt.Errorf("set hallToken err %s", err)
	}
	return nil
}

func Set(key string, value interface{}) error {
	return rc.Set(key, value, 0).Err()
}

func Get(key string) (interface{}, error) {
	val, err := rc.Get(key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key:%s does not exists", key)
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func GetInt(key string) (n int) {
	id, err := Get(key)
	if err != nil {
		log.Printf("[ERROR] get int err: %s", err)
		return
	}
	switch id.(type) {
	case string:
		n = common.StrToInt(id.(string))
	case int:
		n = id.(int)
	}
	return
}
