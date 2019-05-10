package main

// 初始化

import (
	"flag"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	setFlags()
	validateFlags()
	setWebSite()
}

//设置命令行参数
func setFlags() {
	flag.StringVar(&weburl, "url", "", "请设置要克隆的网站地址")
	flag.IntVar(&max, "max", 20, "最大并发数")
	flag.StringVar(&dir, "dir", ".", "下载文件存放目录")
	flag.Parse()
}

//验证命令行参数
func validateFlags() {
	if !strings.HasPrefix(weburl, "http://") && !strings.HasPrefix(weburl, "https://") {
		log.Fatal("url 请以http://或https://开头")
	}

	if max < 1 {
		max = 1
	}

	if _, err := os.Stat(dir); err != nil {
		log.Fatal("该文件夹不存在", err)
	}
}

func getDomain(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return u.Host, err
}

func setWebSite() {
	domain, err := getDomain(weburl)
	if err != nil {
		log.Fatal(err)
	}
	webSite.domain = domain
	webSite.path = filepath.Join(dir, strings.Replace(domain, ".", "_", -1))
}
