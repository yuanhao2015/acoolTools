package test

import (
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"testing"
	"time"
)

func TestConvertTest(t *testing.T) {
	calendar := acoolTools.ConvertUtils.GregorianToLunarCalendar(2020, 2, 1)
	fmt.Println(calendar)
	gregorian := acoolTools.ConvertUtils.LunarToGregorian(calendar[0], calendar[1], calendar[2], false)
	fmt.Println(gregorian)
	days := acoolTools.ConvertUtils.GetLunarYearDays(2021)
	fmt.Println(days)
}

func TestDate(t *testing.T) {
	now := time.Now()
	sub := acoolTools.DateUtil.MinuteAddOrSub(now, -12)
	fmt.Println(sub)
	sub = acoolTools.DateUtil.MinuteAddOrSub(now, 12)
	fmt.Println(sub)
	orSub := acoolTools.DateUtil.HourAddOrSub(now, 12)
	fmt.Println(orSub)
	orSub = acoolTools.DateUtil.HourAddOrSub(now, -12)
	fmt.Println(orSub)
	duration := sub.Sub(orSub)
	fmt.Println(duration)
}
