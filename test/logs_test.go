package test

import (
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestLogs(t *testing.T) {
	acoolTools.Logs.ErrorLog().Println("Error日志测试")
	acoolTools.Logs.InfoLog().Println("Info日志测试")
	acoolTools.Logs.DebugLog().Println("Debug日志测试")
}
