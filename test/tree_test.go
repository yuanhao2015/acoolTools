package test

import (
	"encoding/json"
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"github.com/yuanhao2015/acoolTools/tree"
	"testing"
)

// 定义我们自己的菜单对象
type SystemMenu struct {
	Id    int    `json:"id"`    //id
	Pid   int    `json:"pid"`   //上级菜单id
	Name  string `json:"name"`  //菜单名
	Route string `json:"route"` //页面路径
	Icon  string `json:"icon"`  //图标路径
}

// region 实现ITree 所有接口
func (s SystemMenu) GetTitle() string {
	return s.Name
}

func (s SystemMenu) GetId() int {
	return s.Id
}

func (s SystemMenu) GetPid() int {
	return s.Pid
}

func (s SystemMenu) GetData() interface{} {
	return s
}

func (s SystemMenu) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return s.Pid == 0 || s.Pid == s.Id
}

// endregion

type SystemMenus []SystemMenu

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s SystemMenus) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

func TestGenerateTree(t *testing.T) {
	// 模拟获取数据库中所有菜单，在其它所有的查询中，也是首先将数据库中所有数据查询出来放到数组中，
	// 后面的遍历递归，都在这个 allMenu中进行，而不是在数据库中进行递归查询，减小数据库压力。
	allMenu := []SystemMenu{
		{Id: 1, Pid: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, Pid: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},

		{Id: 3, Pid: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, Pid: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, Pid: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, Pid: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, Pid: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	// 生成完全树
	resp := acoolTools.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), nil)
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	//fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))

	// 模拟选中 '资产' 菜单
	selectedNode := []SystemMenu{allMenu[2]}
	resp = acoolTools.TreeUtils.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), SystemMenus.ConvertToINodeArray(selectedNode))
	bytes, _ = json.Marshal(resp)
	//fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))

	// 模拟从数据库中查询出 '设备'
	device := []SystemMenu{allMenu[5]}
	// 查询 设备 的所有父节点
	respNodes := acoolTools.TreeUtils.FindRelationNode(SystemMenus.ConvertToINodeArray(device), SystemMenus.ConvertToINodeArray(allMenu))
	resp = acoolTools.TreeUtils.GenerateTree(respNodes, nil)
	bytes, _ = json.Marshal(resp)
	fmt.Println(string(acoolTools.PrettyUtils.Color(acoolTools.PrettyUtils.PrettyOptions(bytes, acoolTools.PrettyUtils.GetDefaultOptions()), nil)))
}
