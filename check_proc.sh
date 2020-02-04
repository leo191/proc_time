#!/bin/bash

getpids=$( pgrep $1 )

# sleep 5
allpids=$( ps -eo pid )
for pid in ${getpids[@]}
do
    if [[ $allpids == *"$pid"* ]]; then
        echo "$pid exists"
    else
        echo "$pid dead"
    fi
done