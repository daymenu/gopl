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
	github.com/go-sql-driver/mysql v1.4.1
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
)
