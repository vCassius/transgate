/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 10:47 上午
 * @Desc: 加载配置
 */

package logic

import (
	"TransGate/global"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
	return true
}

func LoadAllConfig() {
	chkGateConf := Exists("./config.json")
	chkBanConf := Exists("./blacklist.json")
	chkWLConf := Exists("./whitelist.json")
	if chkGateConf && chkBanConf && chkWLConf {
		LoadTransGateConf()
		LoadBanConf()
		LoadWLConf()
	} else {
		//NoticeInfo("load conf error!")
		global.Logger.Error("load conf error!")
		os.Exit(1)
	}

}

// LoadTransGateConf 加载网关Json配置
func LoadTransGateConf() {
	JsonParse := NewJsonStruct()
	conf := global.TransGateConf{}
	JsonParse.Load("./config.json", &conf)
	global.IP = conf.IP
	global.Port = conf.Port
	global.MaxUser = conf.MaxUser
	global.DesSrvIP = conf.DesSrvIP
	global.DesSrvPort = conf.DesSrvPort
	global.RunModel = conf.RunModel
}

// LoadBanConf 加载黑名单配置
func LoadBanConf() {
	JsonParse := NewJsonStruct()
	conf := BanConf{}
	JsonParse.Load("./blacklist.json", &conf)
	global.BanIPList = conf.BanIPList
	global.BanIPArr = strings.Split(global.BanIPList, ",") //以,为分隔符，分割字符串导入到字符串数组，用于存储拒绝连接列表
}

// LoadWLConf 加载黑名单配置
func LoadWLConf() {
	JsonParse := NewJsonStruct()
	conf := WLConf{}
	JsonParse.Load("./whitelist.json", &conf)
	global.AcceptIPList = conf.AcceptIPList
	global.AcceptIPArr = strings.Split(global.AcceptIPList, ",") //以,为分隔符，分割字符串导入到字符串数组，用于存储拒绝连接列表
}

type BanConf struct {
	BanIPList string
}
type WLConf struct {
	AcceptIPList string
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

type JsonStruct struct {
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
