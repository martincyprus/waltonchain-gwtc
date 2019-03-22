#!/bin/sh
if [ ! -d "./data" ]; then
	./bin/gwtc_v1.0.0stable --datadir ./data/ init ./settings/wtc_v1.0.0stable.json
fi
if [ "$1" = "--mine" ]; then
./bin/gwtc_v1.0.0stable --networkid 2882 --datadir ./data/ --identity "wtc" $1 --etherbase $2 console
else
./bin/gwtc_v1.0.0stable --networkid 2882 --datadir ./data/ --identity "wtc" console
fi
