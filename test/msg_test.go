/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: msg_test
* @Date: 2021-9-26 17:04
 */
package test

import (
	"github.com/yuanhao2015/acoolTools"
	"testing"
)

func TestMsg(t *testing.T) {

	t.Log(acoolTools.GetMsgUtils.GetMsg(200))
}
