#!/bin/bash

./run.sh process50 > tmp.log 

diff tmp.log ~/tmp/run.gold.log
RESULT=$?
echo -e "\nTRACER result: $RESULT"

