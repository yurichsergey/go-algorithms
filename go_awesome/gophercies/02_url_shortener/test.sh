#!/bin/bash

if lsof -i :8080 > /dev/null 2>&1; then
    echo "Error: port 8080 is already in use"
    exit 1
fi

trap 'lsof -ti :8080 | xargs kill 2>/dev/null' EXIT

go run main/main.go &

until curl -s http://localhost:8080/ > /dev/null 2>&1; do sleep 0.1; done

curl -v http://localhost:8080/urlshort 2>&1 | grep -E "< HTTP|Location"
curl -v http://localhost:8080/urlshort-godoc 2>&1 | grep -E "< HTTP|Location"
curl -s http://localhost:8080/
