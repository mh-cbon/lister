#!/bin/sh

set -ex

rm `which lister`
go install


lister - demo/Tomate:SliceTomate | grep "Empty(" || exit 1;
lister - string:StringSlice | grep "Empty(" || exit 1;

lister - demo/Tomate:SliceTomate | grep "package main" || exit 1;
lister -p nop - demo/Tomate:SliceTomate | grep "package nop" || exit 1;
lister - string:StringSlice | grep "package main" || exit 1;
lister -p nop - string:StringSlice | grep "package nop" || exit 1;

rm -fr gen_test
lister string:gen_test/StringSlice || exit 1;
ls -al gen_test | grep "stringslice.go" || exit 1;
cat gen_test/stringslice.go | grep "Empty(" || exit 1;
cat gen_test/stringslice.go | grep "package gen_test" || exit 1;
rm -fr gen_test

rm -fr demo/poireaux.go demo/tomates.go
go generate demo/main.go
ls -al demo | grep "tomates.go" || exit 1;
cat demo/tomates.go | grep "package main" || exit 1;
cat demo/poireaux.go | grep "package main" || exit 1;
go run demo/*.go | grep "Red" || exit 1;
# rm -fr demo/poireaux.go demo/tomates.go # keep it for demo

rm -fr demo/poireaux.go demo/tomates.go
go generate github.com/mh-cbon/lister/demo
ls -al demo | grep "tomates.go" || exit 1;
cat demo/tomates.go | grep "package main" || exit 1;
cat demo/poireaux.go | grep "package main" || exit 1;
go run demo/*.go | grep "Red" || exit 1;
# rm -fr demo/poireaux.go demo/tomates.go # keep it for demo

go test


echo ""
echo "ALL GOOD!"
