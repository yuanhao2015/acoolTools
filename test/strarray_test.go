package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"reflect"
	"testing"
)

func TestStringToInt64(t *testing.T) {
	//字符串数组转int64
	strings := []string{"1", "23123", "232323"}
	fmt.Println(reflect.TypeOf(strings[0]))
	toInt64, err := acoolTools.StrArrayUtils.StringToInt64(strings)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(toInt64[0]))
}

func TestStringToInt32(t *testing.T) {
	//字符串数组转int64
	strings := []string{"1", "23123", "232323"}
	fmt.Println(reflect.TypeOf(strings[0]))
	toInt64, err := acoolTools.StrArrayUtils.StringToInt32(strings)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(toInt64[0]))
}

func TestArrayDuplication(t *testing.T) {
	//string数组去重
	strings := []string{"hello", "word", "acoolTools", "word"}
	fmt.Println("去重前----------------->", strings)
	duplication := acoolTools.StrArrayUtils.ArrayDuplication(strings)
	fmt.Println("去重后----------------->", duplication)
}
