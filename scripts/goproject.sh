#!/bin/bash
kill -9 $(lsof -t -i :3000) 2>/dev/null

if [ "$1" == 'run' ]; then
    go run main.go &
    pid=$!

    echo "Starting Server"

    while ! ps -p $pid > /dev/null; do
        sleep 2
    done
    echo "Server started and is running"

elif [ "$1" == 'test' ]; then
    echo "Running tests after starting the server"

    go run main.go &
    pid=$!

    echo "Starting Server"

    while ! ps -p $pid > /dev/null; do
        sleep 2
    done
    echo "Server started and is running"
    go clean -testcache
    go test ./... -v
    if [ $? -eq 0 ]; then
        echo "Tests passed"
        kill $pid 2>/dev/null
    else
        echo "Tests failed"
        kill $pid 2>/dev/null
        exit 1
    fi

fi
