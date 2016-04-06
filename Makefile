

mping: mping.go
	go get github.com/mgutz/ansi
	go build -ldflags "-s" mping.go
	strip mping
