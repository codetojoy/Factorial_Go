#!/bin/bash

./run.sh process50 > tmp.log 

diff tmp.log run.gold.log
RESULT=$?
echo -e "\nTRACER result: $RESULT"

