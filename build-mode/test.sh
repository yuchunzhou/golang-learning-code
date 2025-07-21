#! /bin/bash

rm dynlib.* test
go build -buildmode=c-shared -o dynlib.so main.go
gcc -o test test.c dynlib.so
chmod +x test
export LD_LIBRARY_PATH=.
./test
