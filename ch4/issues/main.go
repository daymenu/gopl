package main

import (
	"strings"
)

// /repos/daymenu/phpcap/issues
func main() {
	// in := bufio.NewScanner(os.Stdin)
	// fmt.Println("请你输入git的用户名")
	// if in.Scan() {
	// 	UserName = in.Text()
	// }
	// fmt.Println("请你输入你要操作的仓库名称")
	// if in.Scan() {
	// 	Repos = in.Text()
	// }

	IssuesURL = strings.Join([]string{Domain, "repos", UserName, Repos, "issues"}, "/")
	// issuesList()
	// issues("1")
}
