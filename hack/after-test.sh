#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=./output

mkdir -p $OUTPUT/$PLATFORM
for service_name in `ls $(pwd)/cmd`; do
    kill -9 `pidof $service_name`
done

kill -9 `ps -aux |grep test_mock_mininpool_mw |grep -v grep |awk '{print $2}'|xargs`
