package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
)

func TestClient(t *testing.T) {

	dial, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("主机不存在!!!")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for true {
		content, _ := reader.ReadString('\n')
		dial.Write([]byte(content))
	}

}
