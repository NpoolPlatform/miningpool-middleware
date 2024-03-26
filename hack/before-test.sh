#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../

PLATFORM=linux/amd64
OUTPUT=./output

go run $ROOT_PATH/mock/main.go test_mock_mininpool_mw &

export F2POOL_API="http://127.0.0.1:20180"

mkdir -p $OUTPUT/$PLATFORM
for service_name in `ls $(pwd)/cmd`; do
    cp $(pwd)/cmd/$service_name/*.viper.yaml $OUTPUT/$PLATFORM
    cd $OUTPUT/$PLATFORM; ./$service_name run | grep error &
done

sleep 60
