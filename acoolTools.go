package acoolTools

import (
	"github.com/yuanhao2015/acoolTools/array"
	"github.com/yuanhao2015/acoolTools/bcrypt"
	"github.com/yuanhao2015/acoolTools/captcha"
	"github.com/yuanhao2015/acoolTools/compression"
	"github.com/yuanhao2015/acoolTools/convert"
	"github.com/yuanhao2015/acoolTools/date"
	"github.com/yuanhao2015/acoolTools/desensitized"
	"github.com/yuanhao2015/acoolTools/get_ip"
	"github.com/yuanhao2015/acoolTools/id_utils"
	"github.com/yuanhao2015/acoolTools/logs"
	"github.com/yuanhao2015/acoolTools/openfile"
	"github.com/yuanhao2015/acoolTools/page"
	"github.com/yuanhao2015/acoolTools/pretty"
	"github.com/yuanhao2015/acoolTools/request"
	"github.com/yuanhao2015/acoolTools/response"
	"github.com/yuanhao2015/acoolTools/str"
	"github.com/yuanhao2015/acoolTools/tree"
)

var (
	TableRespUtils    response.TableResp        //通用分页数据返回json
	ApiRespUtils      response.ApiResp          //通用api数据json
	ClientIPUtils     get_ip.GetIpUtil          //ip 获取工具
	StrArrayUtils     array.StrArray            //String 数据工具声明
	Logs              logs.Logs                 //log日志声明
	BcryptUtils       bcrypt.BcryptUtil         //加密工具声明
	DateUtil          date.Date                 //时间操作
	StrUtils          str.StrUtils              //字符串操作
	HttpUtils         request.Request           //http工具
	ConvertUtils      convert.Convert           //公历转农历相关操作
	PageUtils         page.Page                 //分页插件
	IdUtils           id_utils.IdUtils          //id生成工具
	ZipUtils          compression.ZipUtils      //压缩和解压缩工具
	FileUtils         openfile.FileUtils        //IO操作工具
	CaptchaUtils      captcha.Captcha           //验证码工具
	DesensitizedUtils desensitized.Desensitized //敏感数据脱敏
	TreeUtils         tree.Tree                 //菜单树结构化工具
	PrettyUtils       pretty.PrettyUtils        //JSON打印格式化工具
)
