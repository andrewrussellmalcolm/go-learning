go build -o protoc-gen-custom
protoc --plugin=protoc-gen-custom --custom_out=./output testservice.proto