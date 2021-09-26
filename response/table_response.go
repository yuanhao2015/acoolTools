/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: result
* @Date: 2021-9-13 14:46
 */
package response

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanhao2015/acoolTools/e"
	"net/http"
)

// table数据返回
type TableResponse struct {
	Total int         `json:"total"` //总数
	Rows  interface{} `json:"rows"`  //数据
	Code  int         `json:"code"`  //响应编码 200 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
}

// 通用api响应
type TableResp struct {
	t *TableResponse
	c *gin.Context
}

//返回一个成功的消息体
func BuildTable(c *gin.Context, total int, rows interface{}) *TableResp {
	msg := TableResponse{
		Code:  e.SUCCESS,
		Msg:   e.GetMsg(e.SUCCESS),
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		t: &msg,
		c: c,
	}
	return &a
}

//输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}
