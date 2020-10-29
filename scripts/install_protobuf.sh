#!/bin/bash	

PROTOBUF_VERSION=3.13.0	
PROTOC_DOWNLOAD_FILENAME=protoc-$PROTOBUF_VERSION.zip	
PROTOC_DOWNLOAD_DIR=/tmp/protocDownload	

UNAME_S=`uname -s`	
# Only automatically install protoc if on Linux	
if [ "Linux" == "$UNAME_S" ]; then	
    mkdir -p $PROTOC_DOWNLOAD_DIR	

    cd $PROTOC_DOWNLOAD_DIR && \
    curl -L https://github.com/google/protobuf/releases/download/v{$PROTOBUF_VERSION}/protoc-{$PROTOBUF_VERSION}-linux-x86_64.zip -o $PROTOC_DOWNLOAD_FILENAME && \
    unzip $PROTOC_DOWNLOAD_FILENAME && \
    sudo cp ./bin/protoc /usr/local/bin/. && \
    sudo cp -r ./include /usr/local/. && \
    sudo chmod 755 /usr/local/bin/protoc && \
    sudo chmod -R 755 /usr/local/include/google && \
    protoc --version	

    rm -rf $PROTOC_DOWNLOAD_DIR	
else	
    echo "Please manually install protoc based on your OS from https://github.com/protocolbuffers/protobuf/releases"	
fi 