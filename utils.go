package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func guid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func randToken() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 18)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("%x", b)
}

func randUserName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100000)
	return fmt.Sprintf("用户%d", n)
}

func strToInt(str string) int {
	var i int
	i, err = strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}
