#!/bin/sh

set -ex

rm `which lister`
go install

rm -fr gen_test
lister string:gen_test/StringSlice || exit 1;
ls -al gen_test | grep "stringslice.go" || exit 1;
cat gen_test/stringslice.go | grep "Empty(" || exit 1;

rm -fr demo/gen
go generate demo/lib.go
go test
ls -al demo | grep "gen" || exit 1;
ls -al demo/gen | grep "tomates.go" || exit 1;

rm -fr gen_test

echo ""
echo "ALL GOOD!"
