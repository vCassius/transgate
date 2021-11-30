/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 11:31 上午
 * @Desc: 数据交换拷贝
 */

package logic

import (
	"TransGate/global"
	"io"
	"net"
	"sync/atomic"
)

func DataExHandle(proxy net.Conn, target net.Conn) {
	defer func(proxy net.Conn) {
		err := proxy.Close()
		if err != nil {

		}
	}(proxy)
	defer func(target net.Conn) {
		err := target.Close()
		if err != nil {

		}
	}(target)
	ExitChan := make(chan bool, 1)
	go func(proxy net.Conn, target net.Conn, ExChan chan bool) {
		_, _ = io.Copy(proxy, target)
		ExitChan <- true
	}(proxy, target, ExitChan)
	go func(proxy net.Conn, target net.Conn, ExChan chan bool) {
		_, _ = io.Copy(target, proxy)
		ExitChan <- true
	}(proxy, target, ExitChan)

	<-ExitChan
	atomic.AddInt64(&global.ConnectCount, -1)
}
