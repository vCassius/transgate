/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 10:47 上午
 * @Desc:
 */

package logic

import (
	"TransGate/global"
	"TransGate/version"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
)

func StartTransGate() {
	version.PrintVer()
	LoadAllConfig()
	go NoticeCount() //Vincent 20190513 计时函数必须要新开一个协程去做处理,否则会阻塞Socket的
	SourceSrvInfo := global.IP + ":" + global.Port
	server, err := net.Listen("tcp", SourceSrvInfo)

	if err != nil {
		msg := fmt.Sprintf("Unable To Bind IP & Port Error: %s", err.Error())
		global.Logger.Error(msg)
		os.Exit(1) //如果无法监听链接直接退出程序
	}

	msg := fmt.Sprint("TransGate Started ...")
	runEnvInfo := fmt.Sprintf("BindIP: %s | BindPort: %s | MaxUser: %v | DesSrvIP: %s | DesSrvPort: %s | RunModel: %s | CPUCoreMax: %v", global.IP, global.Port, global.MaxUser, global.DesSrvIP, global.DesSrvPort, global.RunModel, global.CPUCoreMax)
	global.Logger.Info(msg)
	global.Logger.Info(runEnvInfo)

	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {

		}
	}(server)
	for {
		switch global.RunModel {
		//黑名单方式运行
		case "bl":
			proxyConn, err := server.Accept()
			isExist := false
			for _, v := range global.BanIPArr {
				banip := proxyConn.RemoteAddr().String()
				banip = banip[0:strings.LastIndex(banip, ":")]
				if v == banip {
					isExist = true
				}
			}
			if isExist {
				global.Logger.Error(fmt.Sprintf("IP Address In BanList, Kick Off IP Address: %s", proxyConn.RemoteAddr()))
				_ = proxyConn.Close()
				continue
			} else {
				//无法接受一个请求
				if err != nil {
					msg := fmt.Sprintf("Unable to accept a request, error: %s", err.Error())
					//NoticeInfo(msg)
					global.Logger.Error(msg)
					continue
				}
				global.Logger.Info(fmt.Sprintf("Accept New Connection: %s", proxyConn.RemoteAddr().String()))
				//转发地址
				targetAddr := fmt.Sprintf("%s:%s", global.DesSrvIP, global.DesSrvPort)
				//net.Dial 是主动拨号链接到别的地址，而不是像listener一样本地建立一个socket监听
				targetConn, err := net.Dial("tcp", targetAddr)
				//拨号链接到转发服务器出错
				if err != nil {
					global.Logger.Error(fmt.Sprintf("net.Dial Unable To Connect To Target: %s, error: %s", targetAddr, err.Error()))
					_ = proxyConn.Close()
					continue
				}
				//判断以下是否达到了配置设定的服务器最大链接数
				if global.ConnectCount >= global.MaxUser {
					global.Logger.Error(fmt.Sprintf("Over Max Connections, Kick Off Connection %s MaxValue --> %v", proxyConn.RemoteAddr().String(), global.MaxUser))
					_ = proxyConn.Close()
					continue
				}
				atomic.AddInt64(&global.ConnectCount, 1) //原子钟加1给在现数
				runtime.GOMAXPROCS(global.CPUCoreMax)    //CPU多核利用
				go DataExHandle(proxyConn, targetConn)
			}
			//白名单方式运行
		case "wl":
			proxyConn, err := server.Accept()
			isExist := false
			for _, v := range global.AcceptIPArr {
				acceptIp := proxyConn.RemoteAddr().String()
				acceptIp = acceptIp[0:strings.LastIndex(acceptIp, ":")]
				if v == acceptIp {
					isExist = true
				}
			}
			if isExist {
				//无法接受一个请求
				if err != nil {
					global.Logger.Error(fmt.Sprintf("Unable to accept a request, error: %s", err.Error()))
					continue
				}
				global.Logger.Info(fmt.Sprintf("Accept New Connection: %s", proxyConn.RemoteAddr().String()))
				//转发地址
				targetAddr := fmt.Sprintf("%s:%s", global.DesSrvIP, global.DesSrvPort)
				//net.Dial 是主动拨号链接到别的地址，而不是像listener一样本地建立一个socket监听
				targetConn, err := net.Dial("tcp", targetAddr)
				//拨号链接到转发服务器出错
				if err != nil {
					global.Logger.Error(fmt.Sprintf("net.Dial Unable To Connect To Target: %s, error: %s", targetAddr, err.Error()))
					err := proxyConn.Close()
					if err != nil {
						return
					}
					continue
				}
				//判断以下是否达到了配置设定的服务器最大链接数
				if global.ConnectCount >= global.MaxUser {
					global.Logger.Error(fmt.Sprintf("Over Max Connections, Kick Off Connection %s MaxValue --> %v", proxyConn.RemoteAddr().String(), global.MaxUser))
					err := proxyConn.Close()
					if err != nil {
						return
					}
					continue
				}
				atomic.AddInt64(&global.ConnectCount, 1) //原子钟加1给在现数
				runtime.GOMAXPROCS(global.CPUCoreMax)    //CPU多核利用
				go DataExHandle(proxyConn, targetConn)
			} else {
				global.Logger.Error(fmt.Sprintf("IP Address Does Not In WhiteList, Kick Off IP Address: %s", proxyConn.RemoteAddr()))
				err := proxyConn.Close()
				if err != nil {
					return
				}
				continue
			}
		default:
			global.Logger.Error(fmt.Sprintf("Wrong Parameter With RunModel: %s", global.RunModel))
			os.Exit(1)
		}
	}
}
