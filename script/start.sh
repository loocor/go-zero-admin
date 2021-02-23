#!/bin/bash
## Author by loocor
## Email loocor@gmail.com

echo "Start sys rpc service @ 8080" &
go run service/sys/sys.go &

echo "Start ums rpc service @ 8081" &
go run service/ums/ums.go &

echo "Start admin api service @ 8888" &
go run api/admin.go &