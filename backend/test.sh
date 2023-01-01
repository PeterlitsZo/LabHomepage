#!/bin/bash

echo "test start"
echo
status_code_position=1
status_code_len=3

echo "test middleware AuthMiddleware"
cmd="http -v localhost:8080/api/v1/users"
status=`$cmd | grep HTTP/ | tee /dev/tyy`
status_code=${status:14:3}
echo
echo "cmd:$cmd"
echo "status is $status"
echo "status code is $status_code"
echo
