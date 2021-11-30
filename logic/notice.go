/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 11:01 上午
 * @Desc: 10分钟打印一下当前连接数
 */

package logic

import (
	"TransGate/global"
	"fmt"
	"time"
)

func NoticeCount() {
	for {
		time.Sleep(10 * time.Minute)
		msg := fmt.Sprintf("Online Connections: %v", global.ConnectCount)
		global.Logger.Info(msg)
	}
}
