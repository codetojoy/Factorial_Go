#!/bin/bash

set -e

export GOPATH=$PWD

HOME_DIR=$PWD
ROOT_DIR=$PWD/src/github.com/codetojoy

cd $ROOT_DIR/config
go test -test.v
CONFIG_RESULT=$?

cd $ROOT_DIR/util
go test -test.v
UTIL_RESULT=$?

cd $ROOT_DIR/prime
go test -test.v
PRIME_RESULT=$?

cd $ROOT_DIR/factor
go test -test.v
FACTOR_RESULT=$?

cd $ROOT_DIR/factorial
go test -test.v
FACTORIAL_RESULT=$?

cd $ROOT_DIR/worker
go test -test.v
WORKER_RESULT=$?

echo -e "\nTRACER config ${CONFIG_RESULT}\n"
echo -e "TRACER util ${UTIL_RESULT}\n"
echo -e "TRACER prime ${PRIME_RESULT}\n"
echo -e "TRACER factor ${FACTOR_RESULT}\n"
echo -e "TRACER factorial ${FACTORIAL_RESULT}\n"
echo -e "TRACER worker ${WORKER_RESULT}\n"

cd $HOME_DIR
echo "Ready."
