#!/bin/bash
set -eu

PACKAGE_NAME=$(basename $PWD)
TEST_FLAGS='-v -race'

if [[ -n $1 ]]; then
  COVERAGE_SERVICE=$1
fi

case $COVERAGE_SERVICE in
  codecov)
    MODE=atomic
    ;;
  coveralls)
    MODE=count
    ;;
  *)
    echo 'unknown service name'
    exit 1
esac

COVERAGE_OUT=coverage.tmp.out
COVERAGE_RESULT=coverage.out

if [ -f "$COVERAGE_RESULT" ]; then
  rm -f $COVERAGE_RESULT
fi

echo "mode: $MODE" > $COVERAGE_RESULT
for pkg in $(glide nv); do
  go test $TEST_FLAGS -cover -covermode=$MODE -coverprofile=$COVERAGE_OUT $pkg
  if [ -f $COVERAGE_OUT ]; then
    sed -i -e "s/_\/home\/ubuntu/github.com\/zchee/g" $COVERAGE_OUT
    tail -n +2 $COVERAGE_OUT >> $COVERAGE_RESULT
    rm -f $COVERAGE_OUT
  fi
done
