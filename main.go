package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net"
	"os"
	"time"
)

var fileName = "G:\\workspace\\peifang.txt"

func bufioOpera(file *os.File) {
	reader := bufio.NewReader(file)

	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(content)
	}
}

// 不推荐,文件过大会影响效率
func ioutilOpera(file *os.File) {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {

	}
	fmt.Println(string(bytes))
}

func init() {
	db, err := sql.Open("mysql", "root:Chiu123456*@tcp(139.199.225.106:3306)/chat-server")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func init1() {

	file, err := os.Open(fileName)
	if err != nil {

	}
	defer file.Close()

	//bufioOpera(file)
	//ioutilOpera(file)

	//openFile(file)

	countFileContent(file)

	file1, errors := os.Open("d:\\a.txt")

	if errors != nil {

	}
	// Stat()方法用来判断是否存在,不存在会报出 路径错误
	info, _ := file1.Stat()
	fmt.Println(info)
}

func countFileContent(file *os.File) {

	content := make([]byte, 64)

	for true {
		read, err := file.Read(content)
		if err == io.EOF {
			break
		}

		for i := 0; i < read; i++ {
			fmt.Println(content[i])
		}

	}

}

func openFile(file *os.File) {

	bytes := make([]byte, 64)

	for true {
		read, err := file.Read(bytes)

		if err == io.EOF {
			break
		}
		fmt.Print(string(bytes[0:read]))
	}

}
func creatFile(fileName string) {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.WriteString("Hello,Tiger\n")
	}

}
func creatFileByBufio(fileName string) {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("Hello,Ken\n")
	}
	writer.Flush()
}

func main() {

	dial, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("主机不存在!!!")
		return
	}
	defer dial.Close()
	reader := bufio.NewReader(os.Stdin)
	go readMsg(dial)
	for true {
		content, _ := reader.ReadString('\n')
		dial.Write([]byte(content))
	}

}

func readMsg(dial net.Conn) {

	var msg []byte = make([]byte, 1024)

	for true {
		read, err := dial.Read(msg)
		if err != nil {
			return
		}
		fmt.Println("收到的信息: ", string(msg[:read]))
	}

}

func main1() {
	//creatFile("G:\\workspace\\peifang1.txt")
	//creatFileByBufio("G:\\workspace\\peifang2.txt")

	fmt.Println(os.Args)

	file, _ := os.Open("D:\\googleDownload")
	defer file.Close()
	fileInfo, err := file.Stat()

	if err != nil {

	}

	fmt.Println("文件夹? : ", fileInfo.IsDir())

	if fileInfo.IsDir() {

		//dir, _ := os.ReadDir(file.Name())
		//for true {
		//	fmt.Println(dir)
		//}

	}

}
