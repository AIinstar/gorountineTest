package cg

import (
	"ipc"
	"encoding/json"
	"errors"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return nil
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) BroadCast(message string) error {
	m := &Message{Content:message} //构造message 结构体
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}

//CenterClient匿名组合了IpcClient，这样就可以直接在代码中调用IpcClient的功
//能了。