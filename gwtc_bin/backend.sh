#!/bin/sh
if [ ! -d "./data" ]; then
	 ./bin/gwtc_v1.0.0stable --datadir ./data/ init ./settings/wtc_v1.0.0stable.json
fi
if [ "$1" = "--mine" ]; then
	nohup ./bin/gwtc_v1.0.0stable --networkid 2882 --datadir ./data/ --identity "wtc" $1 --etherbase $2 > gwtc.log 2>&1 &
else
	nohup ./bin/gwtc_v1.0.0stable --networkid 2882 --datadir ./data/ --identity "wtc" > gwtc.log 2>&1 &
fi
