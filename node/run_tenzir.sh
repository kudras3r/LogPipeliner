#!/bin/bash

# Run tenzir node and pipeline that creates /v1/logs endpoint with fluentbit.

if ! command -v tenzir-node &> /dev/null
then
    echo "tenzir-node could not be found. Please install it first."
    exit 1
fi

tenzir-node&

if ! command -v tenzir -f makeEndPoint.tql &> /dev/null
then
    echo "cant run pipeline makeEndPoint.tql"
    exit 1
fi

tenzir -f makeEndPoint.tql

