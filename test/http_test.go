package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestGet(t *testing.T) {
	get, err := acoolTools.HttpUtils.Debug(true).Get("http://192.168.1.115:8080/user/list ")
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err)
	}
	s := string(get)
	fmt.Println(s)
}
