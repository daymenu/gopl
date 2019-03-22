package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func issuesList() {
	resp, err := http.Get(IssuesURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("fetch %s faild:\n http status: %s\n", IssuesURL, resp.Status)
		return
	}
	var issuess []Issues
	if err := json.NewDecoder(resp.Body).Decode(&issuess); err != nil {
		fmt.Println(err)
	}

	var op string = "n"
	current := 0
	fmt.Printf("该仓库的issuess有%d条\n", len(issuess))
	for {
		if op == "n" {
			issuess, finish := paging(issuess, 10, current)
			current++
			for _, v := range issuess {
				fmt.Printf("issues[%d]:\n\tState:%s\n\tTitle:%s\n", v.Number, v.State, v.Title)
			}
			if finish {
				break
			}
		} else {
			break
		}
		fmt.Println("输入n查看后10条\n,输入其他字符退出查看")
		in := bufio.NewScanner(os.Stdin)
		if in.Scan() {
			op = in.Text()
		}
	}

}

func paging(r []Issues, perPage, current int) ([]Issues, bool) {
	issuesCount := len(r)
	var start, end int
	if current > issuesCount {
		start = issuesCount - 1
	} else {
		start = current
	}
	if start+perPage >= issuesCount-1 {
		end = issuesCount
	} else {
		end = start + perPage
	}
	var finish bool
	if end >= issuesCount {
		finish = true
	} else {
		finish = false
	}
	return r[start:end], finish
}

func issues(id string) (Issues, error) {
	var issues Issues
	oneUrl := strings.Join([]string{IssuesURL, id}, "/")
	resp, err := http.Get(oneUrl)
	if err != nil {
		return issues, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return issues, fmt.Errorf("%s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return issues, err
	}
	return issues, nil
}
