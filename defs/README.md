navigate to [protobuf release page](https://github.com/protocolbuffers/protobuf/releases) and download the latest release for your platform

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
7z x protoc-3.11.2-linux-x86_64.zip -o/usr/local
```

install protoc-gen-go

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

cd to defs directory and run the following command to generate go files from proto files

```bash
cd defs
protoc *.proto --go_out=. ; cd ..
```