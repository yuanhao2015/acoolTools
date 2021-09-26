/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: msg
* @Date: 2021-9-13 15:13
 */
package e

var MsgFlags = map[int]string{
	SUCCESS:        "成功",      //通用成功 200
	ERROR:          "服务器内部错误", //通用服务端失败 500
	INVALID_PARAMS: "错误请求",    //通用客户端失败 400
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
