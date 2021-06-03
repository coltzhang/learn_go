package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("start")
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	// 不带缓冲
	for {
		n, _ := f.Read(buf) // 一次读取buf长度
		if n == 0 {         // 读完了
			break
		}
		os.Stdout.Write(buf[:n])
		//fmt.Println(string(buf[:n]))
	}

	//// 带缓冲
	//r := bufio.NewReader(f)
	//w := bufio.NewWriter(os.Stdout)
	//defer w.Flush()		// 刷盘
	//for {
	//	n, _ := r.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	w.Write(buf[:n])
	//}

	cmd := exec.Command("/bin/ls", "-l")
	//err := cmd.Run()
	buf, err := cmd.Output()
	os.Stdout.Write(buf)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("xxxxxxxxend")
}
