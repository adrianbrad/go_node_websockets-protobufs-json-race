#!/usr/bin/env bash

> ./results/results.txt

for size in s m b
do
    websockets-server 10000 $size >> ./results/results.txt &
    sleep .3
    node ./ws-client/client.js &
    wait

    go tool pprof -show_from HandlerFunc --pdf $GOBIN/websockets-server $GOPATH/pprofs/cpu.pprof > results/cpuProfile_$size.pdf
done
