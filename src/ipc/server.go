package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name()  string
	Handle(method, params string) *Response
}
type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan  string, 0)
	go func(c chan string) {
		for{
			request := <- c

			if request == "CLOSE" { //关闭该链接
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)

			c <- string(b) //返回结果
		}
		fmt.Println("Session closed .")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return  session
}
/*
可以看出，我们用Server接口确定了之后所要实现的业务服务器的统一接口。因为IPC框架
已经解决了“网络层”通信的问题（这里的网络层用channel代替了），业务服务器的使用者只需
要定义支持的指令，然后进行实现即可。稍后的中央服务器就是一个典型的业务服务器实现。
*/