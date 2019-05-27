package ipc

import "encoding/json"

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string)(resp *Response, err error){
	req := &Request{method,params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	client.conn <- string(b)
	str := <-client.conn  //等待返回值

	var respl Response
	err = json.Unmarshal([]byte(str), &respl)
	resp = &respl
	return
}

func(client *IpcClient)Close(){
	client.conn <- "ClOSE"
}
/*
IpcClient的关键函数就是Call()了，这个函数会将调用信息封装成一个JSON格式的字符
串发送到对应的channel，并等待获取反馈。
*/