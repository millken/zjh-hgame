package common

import (
	"fmt"

	"gopkg.in/redis.v3"
)

type Session struct {
	ghostId, hallToken string
	client             *redis.Client
}

func NewSession(ghostId, hallToken string) (*Session, error) {
	return &Session{
		ghostId:   ghostId,
		hallToken: hallToken,
	}, nil
}

func (s *Session) SetRedis(client *redis.Client) error {
	s.client = client
	return nil
}

func (s *Session) Set(key string, value interface{}) error {
	return s.client.Set(s.Key(key), value, 3600).Err()
}

func (s *Session) Get(key string) (interface{}, error) {
	val, err := s.client.Get(s.Key(key)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key:%s does not exists", s.Key(key))
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *Session) Key(key string) string {
	redis_key := fmt.Sprintf("%s:%s", s.hallToken, key)
	return redis_key
}
