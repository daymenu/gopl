module gopl

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net => github.com/golang/net v0.0.0-20190520210107-018c4d40a106
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190520201301-c432e742b0af
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190520220859-26647e34d3c0
)

require (
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gotk3/gotk3 v0.0.0-20191027191019-60cba67d4ea4
	github.com/lxn/walk v0.0.0-20191113135339-bf589de20b3c
	github.com/lxn/win v0.0.0-20191106123917-121afc750dd3 // indirect
	github.com/mattn/go-gtk v0.0.0-20191030024613-af2e013261f5 // indirect
	github.com/mattn/go-pointer v0.0.0-20190911064623-a0a44394634f // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
)
