#!/bin/bash
example_name=$1
flag=$2
cd $example_name

if [ -f "schema.yaml" ]; then
    if [ -f "schema.json" ]; then
        exit 1
    fi
fi

echo "Building $example_name"
schema -ts flag
echo "compiling "$example_name"_bg.wasm"
call asc ts/"$example_name"/lib.ts --lib d:/work/node_modules --binaryFile ts/pkg/"$example_name"_ts.wasm
rem call asc ts/"$example_name"/lib.ts --lib d:/work/node_modules --binaryFile ts/pkg/"$example_name"_ts.wasm --textFile ts/pkg/"$example_name"_ts.wat
