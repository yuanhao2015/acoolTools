package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestRemoveSuffix(t *testing.T) {
	fullFilename := "test.txt"
	suffix, _ := acoolTools.StrUtils.RemoveSuffix(fullFilename)
	fmt.Println(suffix)
	fullFilename = "/root/home/test.txt"
	suffix, _ = acoolTools.StrUtils.RemoveSuffix(fullFilename)
	fmt.Println(suffix)
}

func TestStringReplacePlaceholder(t *testing.T) {
	s := "你是我的{},我是你的{}"
	placeholder, err := acoolTools.StrUtils.ReplacePlaceholder(s, "唯一", "所有")
	if err == nil {
		fmt.Println(placeholder)
	}
}

func TestGetSuffix(t *testing.T) {
	fullFilename := "test.txt"
	suffix, _ := acoolTools.StrUtils.GetSuffix(fullFilename)
	fmt.Println(suffix)
	fullFilename = "/root/home/test.txt"
	suffix, _ = acoolTools.StrUtils.GetSuffix(fullFilename)
	fmt.Println(suffix)
}

func TestHasStr(t *testing.T) {
	str := ""
	empty := acoolTools.StrUtils.HasEmpty(str)
	fmt.Println(empty)
	str = "11111"
	empty = acoolTools.StrUtils.HasEmpty(str)
	fmt.Println(empty)
}
