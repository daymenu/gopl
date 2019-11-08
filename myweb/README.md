# 最简单的web服务器
```go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8080", nil)
```

# mux 的web服务器
```go
package main

import "net/http"

type db struct{}

func main() {
	mux := http.NewServeMux()
	d := db{}
	mux.Handle("/hello/", &d)
	http.ListenAndServe(":8080", mux)
}

// ServeHttp hello 服务器
func (d *db) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

```

# 使用自定义路由的服务器
```
package main

import "net/http"

type db struct{}

func main() {
	d := db{}
	http.ListenAndServe(":8080", &d)
}

// ServeHttp hello 服务器
func (d *db) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
```