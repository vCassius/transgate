/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 10:36 上午
 * @Desc:
 */

package global

import (
	"TransGate/logs"
	"runtime"
)

// TransGateConf 全局变量来存储配置信息
type TransGateConf struct {
	IP         string
	Port       string
	MaxUser    int64  //允许连接的用户最大值
	DesSrvIP   string //目标服务器OP
	DesSrvPort string //目标服务器端口
	RunModel   string //定义了运行模式，wl代表白名单,bl代表黑名单，白名单模式加载wl.json(只允许白名单访问),黑名单加载bl(除了黑名单以外的ip都可以访问)
}

type BanConf struct {
	BanIPList string
}
type WLConf struct {
	AcceptIPList string
}

type JsonStruct struct {
}

var IP string
var Port string
var MaxUser int64
var DesSrvIP string
var DesSrvPort string
var BanIPList string //设计用于禁止链接的ip列表
var BanIPArr []string
var AcceptIPList string
var AcceptIPArr []string
var Logger = logs.InitLogger("./syslog/TransGateSys.log", "debug")
var ConnectCount int64
var RunModel string               //服务器运行模式
var CPUCoreMax = runtime.NumCPU() //全局变量记录cpu核心数
