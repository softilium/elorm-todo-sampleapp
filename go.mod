module github.com/softilium/elorm-todo-backend

go 1.23.4

require (
	github.com/lib/pq v1.10.9
	github.com/softilium/elorm v0.0.0
	modernc.org/sqlite v1.38.0
)

require (
	github.com/flosch/pongo2/v6 v6.0.0
	github.com/gorilla/handlers v1.5.2
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/sys v0.33.0 // indirect
	modernc.org/libc v1.66.3 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
)

replace github.com/softilium/elorm => ../elorm
