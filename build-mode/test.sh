#! /bin/bash

rm *.so *.h test

go build -buildmode=c-shared -o lib1.so lib1.go
go build -buildmode=c-shared -o lib2.so lib2.go

gcc -o test test.c lib1.so lib2.so

export LD_LIBRARY_PATH=.

chmod +x test
./test
