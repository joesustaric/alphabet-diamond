#!/usr/bin/env bash


if [ "$1" == "test" ]; then
  go test -v .
  exit "$?"
fi

if [ "$1" == "run" ] ||[ "$#" -ge 2 ]; then
  go run main.go -i "$2"
  exit "$?"
fi

echo "Usage"
echo "./go.sh test|run"
