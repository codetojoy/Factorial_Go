#!/bin/bash

export GOPATH=$PWD

HOME_DIR=$PWD
ROOT_DIR=src/github.com/codetojoy

cd $ROOT_DIR/prime
go test -test.v
PRIME_RESULT=$?

cd $ROOT_DIR/factor
go test -test.v
FACTOR_RESULT=$?

cd $ROOT_DIR/factorial
go test -test.v
FACTORIAL_RESULT=$?

echo -e "\nTRACER prime ${PRIME_RESULT}\n"
echo -e "TRACER factor ${FACTOR_RESULT}\n"
echo -e "TRACER factorial ${FACTORIAL_RESULT}\n"

cd $HOME_DIR
echo "Ready."
