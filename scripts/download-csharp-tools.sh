#!/bin/bash

echo "fetching csharp grpc tools"

curl https://www.nuget.org/api/v2/package/Grpc.Tools/ -L -o tools.zip

unzip tools.zip -d csharp/tools

# cleanup
rm tools.zip
find csharp/tools -type f -maxdepth 1 -print0 | xargs -0  -I {} rm {}
find csharp/tools -type d -maxdepth 1 -not -name 'tools' -print0 | xargs -0  -I {} rm -rf {}

os=$(uname)
bit=$(getconf LONG_BIT)
filename=""
if [[ "$os" == 'Darwin' ]]; then
    filename="macosx_x${bit}"
elif [[ "$os" == 'Linux' ]]; then
    if [[ $(arch) == "arm64" ]]; then
        filename="linux_arm64"
    else
        filename="linux_x${bit}"
    fi
else
    filename="windows_x${bit}"
fi

mv csharp/tools/tools/* csharp/tools
rm -rf csharp/tools/tools
mv csharp/tools/$filename/grpc_csharp_plugin tmp
rm -rf csharp/tools/*
mv tmp csharp/tools/grpc_csharp_plugin

# collect
#find csharp/tools -type f -not -name 'protoc' -print0 | xargs -0 -I {} mv csharp/tools/{}