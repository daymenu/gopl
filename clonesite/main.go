package main

import "fmt"

var (
	weburl  string
	max     int
	dir     string
	webSite WebSite
)

type WebSite struct {
	domain string
	path   string
}

func main() {
	fmt.Println(webSite)
	// if err := os.Mkdir(filepath.Join(dir, "site"), os.ModeDir); err != nil {
	// 	log.Fatal("创建文件夹：", err)
	// }

}
