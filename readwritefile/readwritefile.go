package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFileF() {
	file, err := os.Open(`E:\cccc.txt`)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	bytes := make([]byte, 100, 100)
	num, err := file.Read(bytes)
	if err != nil {
		fmt.Println("读文件发生错误", err)
		return
	}
	fmt.Println(string(bytes[:num]))
}

func writeFileF() {
	file, err := os.OpenFile(`E:\bbbb.txt`, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	content := "雨花开在雨天\n"

	num, err := file.Write([]byte(content))
	if err != nil {
		return
	}
	fmt.Printf("成功写入%d字节", num)
}

func readFile2() {
	file, err := os.Open("E:\\cccc.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		readString, err := reader.ReadString('\n') // 碰到换行符暂停
		if err == io.EOF {                         // 读到文件末尾
			fmt.Printf("%s", readString)
			break
		} else if err != nil {
			fmt.Println("读取文件失败", err)
			return
		}
		fmt.Printf("%s", readString)
	}

}

func main() {
	// readFileF()
	// writeFileF()
	readFile2()
}
