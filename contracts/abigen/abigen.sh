#!/bin/bash
sol=${1:-""}
out=${2:-""}

cat $sol | jq .abi > .abi.json
abigen --abi=.abi.json --pkg=$out --out $out.go
rm .abi.json
