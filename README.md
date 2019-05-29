This is a minimal test case to reproduce https://github.com/golang/go/issues/24882.

To reproduce, run
 go get github.com/bemasc/gobind-fail
 gomobile bind -target android github.com/bemasc/gobind-fail/gamma
