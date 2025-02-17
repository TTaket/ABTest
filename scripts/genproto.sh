#!/bin/bash
BaseDir=$(dirname "$0")/..
ProtoDir=$BaseDir/pkgs/proto
ProtosDir=$ProtoDir/protos

# Set the working directory to the root directory
cd $BaseDir 

# Generate gRPC files from proto files in /pkgs/proto and place them in /apps
for proto_file in $(find $ProtosDir -name '*.proto'); 
do
    # Extract the base name of the proto file (without extension)
    proto_base_name=$(basename "$proto_file" .proto)
    
    # Create a directory in apps with the proto base name
    target_dir=$ProtoDir/pb_$proto_base_name
    mkdir -p "$target_dir"
    
    # create link file to the proto file
    if [ ! -f "$target_dir/$proto_base_name.proto" ]; then
        ln "$proto_file" "$target_dir/$proto_base_name.proto"
    fi

    # Generate gRPC files and place them in the target directory
    protoc --go_out=$target_dir --go_opt=paths=source_relative \
    --go-grpc_out=$target_dir --go-grpc_opt=paths=source_relative \
    --proto_path="$ProtosDir" "$proto_file"
done