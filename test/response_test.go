/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: response_test.go
* @Date: 2021-9-26 17:16
 */
package test

import (
	"github.com/yuanhao2015/acoolTools"
	"github.com/yuanhao2015/acoolTools/response"
	"testing"
)

func TestResponse(t *testing.T) {
	//通用api返回json数据
	acoolTools.ApiRespUtils.SetMsg("test").SetCode(200).SetData([]string{
		"1", "2", "3",
	}).WriteJsonExit()
	//通用数据表格返回json数据
	response.BuildTable(nil, 11, "2222")
	acoolTools.TableRespUtils.WriteJsonExit()
}
