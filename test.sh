#!/bin/sh

set -ex

rm `which lister`
go install


lister string:StringSlice | grep "Empty(" || exit 1;

rm -fr gen_test
lister string:gen_test/StringSlice || exit 1;
ls -al gen_test | grep "stringslice.go" || exit 1;
cat gen_test/stringslice.go | grep "Empty(" || exit 1;
rm -fr gen_test

rm -fr demo/gen
go generate demo/lib.go
ls -al demo | grep "gen" || exit 1;
ls -al demo/gen | grep "tomates.go" || exit 1;
# rm -fr demo/gen # keep it for demo

rm -fr demo/gen
go generate github.com/mh-cbon/lister/demo
ls -al demo | grep "gen" || exit 1;
ls -al demo/gen | grep "tomates.go" || exit 1;
# rm -fr demo/gen # keep it for demo

go test


echo ""
echo "ALL GOOD!"
