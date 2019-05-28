package cg

import (
		"fmt"
)

/*
我们接下来该实现中央服务器了。中央服务器为全局唯一实例，从原则上需要承担以下责任：
 在线玩家的状态管理
 服务器管理
 聊天系统
我们现在因为没有实现其他服务器，所以服务器管理这一块先空着。目前聊天系统也先只实
现了广播。要实现房间内聊天或者私聊，其实都可以根据当前的实现进行扩展。代码清单4-8实
现了在线玩家的管理。

*/
type Player struct {
	Name string "name"
	Level int "level"
	Exp int "exp"
	Room int "room"

	mq chan *Message //等待收取消息
}

type Room struct {
	Name string "name"
	Level int  "level"
}



func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"",0,0,0, m}
	go func(p *Player) {
		msg := <-p.mq
		fmt.Println(p.Name, "received message:", msg.Content)
	}(player)
	return player
}