## Initialize
set GO111MODULE=on
go clean
go mod init github.com/anthonydenecheau/csvtools
go mod vendor # if you have vendor/ folder, will automatically integrate
go build

## Run
set GO111MODULE=on
go run main.go readTC -directory C:\developpement\data\temp

go run main.go find -prefix tc -directory C:\developpement\data\temp
go run main.go read -prefix tc -directory C:\developpement\data\temp
go run main.go process -prefix tc -directory C:\developpement\data\temp -out result.csv

## References
https://nicedoc.io/jszwec/csvutil
