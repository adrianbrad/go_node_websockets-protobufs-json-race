#!/usr/bin/env bash

> ./results/results.txt

for size in s m b
do
    go run ./ws-server/websockets-server.go 100000 $size >> ./results/results.txt &
    sleep 1
    node ./ws-client/client.js &
    wait

    go tool pprof -show_from HandlerFunc --pdf $GOBIN/websockets-server $GOPATH/pprofs/cpu.pprof > results/cpuProfile_$size.pdf
done
