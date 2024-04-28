package chat

import (
	"fmt"
	"net"
	"testing"
)

var connMap map[string]net.Conn = make(map[string]net.Conn)

func TestOpenServer(t *testing.T) {

	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("监听失败!!!")
		return
	}

	for true {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败!!!")
			return
		}
		s := conn.RemoteAddr().String()
		connMap[s] = conn
		go chat(conn)
	}

}

func chat(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr.String(), "连接了本服务器")
	var msg []byte = make([]byte, 1024)

	for true {
		read, err := conn.Read(msg)
		if err != nil {
			delete(connMap, remoteAddr.String())
			return
		}

		for _, val := range connMap {
			if val != conn {
				conn.Write(msg[:read])
			}
		}

		fmt.Println(string(msg[:read]))
	}
}
