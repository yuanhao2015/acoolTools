package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestDesensitized(t *testing.T) {
	//mail desensitization
	mail := "testhello@gmail.com"
	star := acoolTools.DesensitizedUtils.DefaultDesensitized(mail)
	fmt.Println("mail-------------------->", star)
	//phone desensitization
	phone := "13333333333"
	desensitized := acoolTools.DesensitizedUtils.DefaultDesensitized(phone)
	fmt.Println("phone------------------->", desensitized)
	//customize desensitization
	//customize := "sadasdasdkljkldfjlkdjflkjsdf"
	//hideStar := acoolTools.DesensitizedUtils.CustomizeHash(customize, 4, 14)
	//fmt.Println("customize--------------->", hideStar)
}
