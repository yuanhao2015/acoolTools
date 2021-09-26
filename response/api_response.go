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

//api数据
type ApiResponse struct {
	Code int         `json:"code"` // 响应编码 200 成功 500 错误 403 无权限
	Msg  string      `json:"msg"`  // 描述
	Data interface{} `json:"data"` // 返回数据
}

//通用api响应
type ApiResp struct {
	r *ApiResponse
	c *gin.Context
}

//返回一个成功的消息体
func SucessResp(c *gin.Context) *ApiResp {
	msg := ApiResponse{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

//返回一个服务端的错误的消息体
func ErrorResp(c *gin.Context) *ApiResp {
	msg := ApiResponse{
		Code: e.ERROR,
		Msg:  e.GetMsg(e.ERROR),
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

//返回一个客户端的错误消息体
func ErrorInvalidResp(c *gin.Context) *ApiResp {
	msg := ApiResponse{
		Code: e.INVALID_PARAMS,
		Msg:  e.GetMsg(e.INVALID_PARAMS),
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

//设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.r.Msg = msg
	return resp
}

//设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.r.Code = code
	return resp
}

//设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.r.Data = data
	return resp
}

//输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.r)
	resp.c.Abort()
}
