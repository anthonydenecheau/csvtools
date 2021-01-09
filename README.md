***Initialize
set GO111MODULE=on
go clean
go mod init github.com/anthonydenecheau/csvtools
go mod vendor # if you have vendor/ folder, will automatically integrate
go build

***Run
set GO111MODULE=on
go run main.go readTC -directory C:\developpement\data\temp

***References
https://nicedoc.io/jszwec/csvutil
