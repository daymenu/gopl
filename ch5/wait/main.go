package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/**
处理错误的5中策略
1. 向上传递错误
2. 重试
3. 输出错误信息并报错，一般只在main函数中进行
4. 输出错误信息
5. 忽略
*/
func main() {
	url := "http://www.baidu.com"
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
		os.Exit(1)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Panicf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
