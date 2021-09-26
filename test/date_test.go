package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
	"time"
)

func TestFormatToString(t *testing.T) {
	now := acoolTools.DateUtil.Now()
	toString := acoolTools.DateUtil.FormatToString(now, "YYYY-MM-DD hh:mm:ss")
	fmt.Println(toString)
	toString = acoolTools.DateUtil.FormatToString(now, "YYYYMMDD hhmmss")
	fmt.Println(toString)
}

func TestDate_UnixToTime(t *testing.T) {
	fmt.Println(acoolTools.DateUtil.UnixToTime(acoolTools.DateUtil.Now().Unix()))
}

func TestDate_IsZero(t *testing.T) {
	t2 := time.Time{}
	zero := acoolTools.DateUtil.IsZero(t2)
	fmt.Println(zero)
	zero = acoolTools.DateUtil.IsZero(acoolTools.DateUtil.Now())
	fmt.Println(zero)
}

func TestInterpretStringToTimestamp(t *testing.T) {
	timestamp, err := acoolTools.DateUtil.InterpretStringToTimestamp("2021-05-04 15:12:59", "YYYY-MM-DD hh:mm:ss")
	if err != nil {
		acoolTools.Logs.ErrorLog().Println(err.Error())
	}
	fmt.Println(timestamp)
}

func TestUnixToTime(t *testing.T) {
	unix := acoolTools.DateUtil.Now().Unix()
	fmt.Println("时间戳----------------------->", unix)
	toTime := acoolTools.DateUtil.UnixToTime(unix)
	fmt.Println(toTime)
}

func TestGetWeekDay(t *testing.T) {
	now := acoolTools.DateUtil.Now()
	day := acoolTools.DateUtil.GetWeekDay(now)
	fmt.Println("今天是-----------------周", day)
}

//时间计算
func TestTimeAddOrSub(t *testing.T) {
	now := acoolTools.DateUtil.Now()
	fmt.Println("现在时间是--------------------->", now)
	sub := acoolTools.DateUtil.MinuteAddOrSub(now, 10)
	fmt.Println("分钟计算结果-------------------->", sub)
	sub = acoolTools.DateUtil.MinuteAddOrSub(now, -10)
	fmt.Println("分钟计算结果-------------------->", sub)
	sub = acoolTools.DateUtil.HourAddOrSub(now, 10)
	fmt.Println("小时计算结果-------------------->", sub)
	sub = acoolTools.DateUtil.HourAddOrSub(now, -10)
	fmt.Println("小时计算结果-------------------->", sub)
	sub = acoolTools.DateUtil.DayAddOrSub(now, 10)
	fmt.Println("天计算结果-------------------->", sub)
	sub = acoolTools.DateUtil.DayAddOrSub(now, -10)
	fmt.Println("天计算结果-------------------->", sub)
}
