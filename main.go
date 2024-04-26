package main

import (
	"github.com/starskim/DDBOT"
	_ "github.com/starskim/DDBOT-Plugins/acfun"
	_ "github.com/starskim/DDBOT-Plugins/bilibili"
	_ "github.com/starskim/DDBOT-Plugins/douyu"
	_ "github.com/starskim/DDBOT-Plugins/huya"
	_ "github.com/starskim/DDBOT-Plugins/twitcasting"
	_ "github.com/starskim/DDBOT-Plugins/weibo"
	_ "github.com/starskim/DDBOT-Plugins/youtube"
	// 别忘记在这里引入刚刚编写的插件
)

func main() {
	// 使用默认的日志格式配置
	DDBOT.SetUpLog()
	// 启动bot，会自动阻塞
	DDBOT.Run()
}
