#!/bin/bash

echo "working from $PWD" 

rm -rf typescript/{src,dist}

protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=typescript proto/*.proto

pushd typescript > /dev/null || exit 1

# rename proto folder to src
mv -f proto src

pushd src > /dev/null || exit 1

# remove old index and create new one with all ts files generated
# rm index.ts 2> /dev/null
# for e in ./*.ts
# do
#   echo "export * from \"$(echo $e | sed 's/.ts//')\"" >> index.ts
# done

# go back to root
popd > /dev/null
popd > /dev/null

npm run build:typescript
