package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func wrtie(conn net.Conn, txt string) error {

	_, err := conn.Write([]byte(txt))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func input() string {
	fmt.Printf("%v请输入聊天内容🗨️：\n", time.Now().Format("[2006-01-02 15:04:05]"))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return string(scanner.Text())
}

func main() {
	fmt.Printf("%v正在连接『🔞🔞🔞』群聊系统\n", time.Now().Format("[2006-01-02 15:04:05]"))
	conn, err := net.Dial("tcp", "127.0.0.1:8011")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("已连接『🔞🔞🔞』群聊系统")

	fmt.Println("注册 ID ：")
	var userName string
	//使用Scan输入，不允许出现空格
	_, _ = fmt.Scan(&userName)
	_, _ = conn.Write([]byte(userName))

	buf2 := make([]byte, 4096)
	n, err := conn.Read(buf2)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}

	fmt.Println(string(buf2[:n]))
	fmt.Println("📣📣📣提示!!长时间没有发送消息会被系统强制踢出")
	go func() {
		for {
			// buffer1 := make([]byte, 4096)
			//这里使用Stdin标准输入，因为scanf无法识别空格
			info := input()
			if err := wrtie(conn, info); err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} //写操作出现error的概率比较低，这里省去判断
		}
	}()
	for {
		buffer2 := make([]byte, 4096)
		n, err := conn.Read(buffer2)
		if n == 0 {
			fmt.Println("『🔞🔞🔞』群聊系统已关闭,正在退出……")
			return
		}
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Print(string(buffer2[:n]))

	}

}
