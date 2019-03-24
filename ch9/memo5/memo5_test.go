package memo3

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	data := []string{
		"http://www.baidu.com",
		"http://www.baidu.com",
		"http://www.xgo.cn",
		"http://www.xgo.cn",
	}
	var n sync.WaitGroup
	m := New(HttpGetBody)
	for _, url := range data {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Println(err)
			}
			t.Logf("%s %ds %d bytes", url, time.Since(start), len(value.([]byte)))

			n.Done()
		}(url)
	}
	n.Wait()
}

func HttpGetBody(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
