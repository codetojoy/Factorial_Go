#!/bin/bash

cd prime
go test -test.v
PRIME_RESULT=$?
cd ..

cd factor
go test -test.v
FACTOR_RESULT=$?
cd ..

cd factorial
go test -test.v
FACTORIAL_RESULT=$?
cd ..

go test -test.v
ROOT_RESULT=$?

echo -e "\nTRACER prime ${PRIME_RESULT}\n"
echo -e "TRACER factor ${FACTOR_RESULT}\n"
echo -e "TRACER factorial ${FACTORIAL_RESULT}\n"
echo -e "TRACER root ${ROOT_RESULT}\n"
