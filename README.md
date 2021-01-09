set GO111MODULE=on
go mod init github.com/anthonydenecheau/golang-scc
go mod vendor # if you have vendor/ folder, will automatically integrate
go build

set GO111MODULE=on
go run main.go readTC -directory C:\developpement\data\temp

https://nicedoc.io/jszwec/csvutil
