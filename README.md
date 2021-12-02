# TransGate
### 作者信息
Author: Vincent
E-Mail: 46603415@qq.com
QQ: 46603415
### 配置参数
#### 基本配置
`config.json`
```json
{
  "IP":"0.0.0.0",                  
  "Port":"7200",                   
  "DesSrvIP":"192.168.10.249",     
  "DesSrvPort":"3389",             
  "MaxUser":4,                      
  "RunModel":"bl"                   
}
```
```conf
  "IP":"0.0.0.0",                   //转发服务器ip地址
  "Port":"7200",                    //转发服务器端口
  "DesSrvIP":"192.168.10.249",      //目标服务器IP地址
  "DesSrvPort":"3389",              //目标服务器端口
  "MaxUser":4                       //最大连接数
  "RunModel":"bl"                   //运行白名单模式写参数wl，运行黑名单模式写参数bl.
```
#### 白名单
`whitelist.json`
json数组,分隔符为,修改后必须重启网关生效
```json
{
  "AcceptIPList":"127.0.0.1,192.168.10.255"
}
```
#### 黑名单
`blacklist.json`
json数组,分隔符为,修改后必须重启网关生效
```json
{
  "BanIPList":"192.168.10.255,192.168.10.254,127.0.0.1" 
}
```
#### 编译
```shell
git clone https://github.com/Cassuis/transgate.git
cd transgate
go mod tidy
sh build-linux64.sh
sh build-macos.sh
sh build-win64.sh
```
#### 运行
- Linux 
  - 前台运行
    ```shell
    ./TransGate
    ```
  - 后台运行
    ```shell
    ./TransGate &
    ```
- Windows
  双击运行,日志输出在syslog下
- Mac
  同Linux