#!/bin/bash
## Author by loocor
## Email loocor@gmail.com

echo "Start sys rpc service @ 8080" &
go run service/sys/sys.go &

echo "Start ums rpc service @ 8081" &
go run service/ums/ums.go &

echo "Start pms rpc service @ 8082" &
go run service/pms/pms.go &

echo "Start oms rpc service @ 8083" &
go run service/oms/oms.go &

echo "Start sms rpc service @ 8084" &
go run service/sms/sms.go &

echo "Start admin api service @ 8888" &
go run api/admin.go &

echo "Start front api service @ 80" &
go run front/front.go &