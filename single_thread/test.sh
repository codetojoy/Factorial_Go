#!/bin/bash

cd prime
go test -test.v
PRIME_RESULT=$?
cd ..

go test -test.v
ROOT_RESULT=$?

if [ "$PRIME_RESULT" -eq "0" ] 
then
    echo -e "\nTRACER prime ok\n"
else
    echo -e "\nTRACER prime FAIL !\n"
fi

if [ "$ROOT_RESULT" -eq "0" ] 
then
    echo -e "TRACER root ok\n"
else
    echo -e "TRACER root FAIL !\n"
fi
