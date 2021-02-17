#!/bin/bash
## Author by loocor
## Email loocor@gmail.com

echo "Try to kill rpc & api service" &

kill -9 $(lsof -i:8080 -t) &
kill -9 $(lsof -i:8081 -t) &
kill -9 $(lsof -i:8082 -t) &
kill -9 $(lsof -i:8083 -t) &
kill -9 $(lsof -i:8084 -t) &
kill -9 $(lsof -i:8888 -t) &
kill -9 $(lsof -i:80 -t) &
