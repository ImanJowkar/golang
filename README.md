# Golang

### Install on Linux

```sh

wget https://go.dev/dl/go1.25.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.25.3.linux-amd64.tar.gz

vim /etc/bash.bashrc
-----
# add at the end of this file
export PATH=$PATH:/usr/local/go/bin
-----

go version



go mod init github.com/myuser/golang
go run main.go

go build main.go
go build 


set GOOS=windows
set GOARCH=amd64
go build -o winapp
---------

set GOOS=linux
set GOARCH=amd64
go build -o linuxapp


go fmt
go env
go env GOPATH

```