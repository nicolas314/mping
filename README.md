# mping
Ping multiple hosts and report which ones are up/down or DNS-unknown
Useful to determine which hosts are alive on a local network.
Dependency on github.com/mgutz/ansi for ANSI terminal colors.

Build:

    go get github.com/mgutz/ansi
    go build mping.go

Alternatively, use 'make' to build without debug symbols and stripped
(requires Go 1.6).

MIT License. Enjoy!

