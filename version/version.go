/**
 * @Author: Vincent
 * @Author: 46603415@qq.com
 * @Date: 2020/9/17 11:42 上午
 * @Desc: 打印版本信息
 */

package version

import (
	"TransGate/global"
	"fmt"
	"github.com/gogf/gf"
)

var (
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
	Author         = "Vincent"
	Email          = "46603415@qq.com"
)

func PrintVer() {
	msgVer := fmt.Sprintf("Author: %s| E-Mail: %s | Build Data: %s | Golang Version: %s | GF Version: %s", Author, Email, BuildTime, BuildGoVersion, gf.VERSION)
	global.Logger.Info(msgVer)
	//fmt.Print(msgVer)
}
