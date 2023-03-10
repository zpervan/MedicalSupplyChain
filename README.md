[![Ubuntu Client](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/ubuntu_client.yml/badge.svg)](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/ubuntu_client.yml)
[![Ubuntu Server](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/ubuntu_server.yml/badge.svg)](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/ubuntu_server.yml)
[![Windows Client](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/windows_client.yml/badge.svg)](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/windows_client.yml)
[![Windows Server](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/windows_server.yml/badge.svg)](https://github.com/zpervan/MedicalSupplyChain/actions/workflows/windows_server.yml)

# MedicalSupplyChain #

## Environment ##
- [Rust 1.67](https://www.rust-lang.org/tools/install)
- [Golang 1.20](https://go.dev/doc/install)

## Setup ##

### Ubuntu ###

Besides the above-mentioned prerequisites, also install the following dependencies:
```shell
sudo apt install build-essential autoconf libtool pkg-config
```

You can install the `protoc` compiler in two ways:
1. Download [protoc 21.12](https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip), build it then add the built binary (`bin` file) to your `PATH` environment
2. Install with `sudo apt-get install protobuf-compiler`

I recommend to install it via the first step as the `apt-get` repositories tend to lag behind, but it's up to you. Also, don't forget to generate new proto Go files if you have a different version than 21.6. 

Install the Rust `protobuf` extensions by executing:
```shell
cargo install grpc-compiler
```

Navigate to the `server` folder and execute the following command:
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go mod tidy
```

Position your terminal at the root of the project and execute the following command to generate `protobuf` and `protobuf-grpc` files for the `server`:
```shell
protoc --proto_path=./messages --go_out=./server/proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=./server/proto medical_supplies.proto
```

Explanation:
- `proto_path`    - directory path where the `.proto` files are located
- `go_out`        - directory path where the Golang `proto` files will be generated; contains the `message` data (structs)
- `go_opt`        - a flag which tells `protoc` to generates the Golang files inside the `proto` folder, otherwise it will create folders as specified in `proto_path`
- `go-grpc_out`   - directory path where the Golang `proto-grpc` files will be generated; contains the `service` data
- `go-grpc_opt`   - a flag which tells `protoc` to generates the Golang grpc files inside the `proto` folder, otherwise it will create folders as specified in `proto_path`

The `client` has a `build.rs` script which is generating the files each time a build is executed.

### Windows ###

Download the `protoc-21.12` protobuf compiler from this [link](https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-win64.zip).

In your terminal, navigate to the `client` folder and execute:
```shell
cargo install grpc-compiler
```

Navigate to the `server` folder and execute:
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest # Installs the protoc compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  # Installs the protoc plugin for Go
go mod tidy
```

## Run ##

Open two separate terminal windows. In one, navigate to the `server` folder and run:
```shell
go run .
```

In the other window, navigate to the `client` folder and run:
```shell
cargo run
```

## Additional ##
- [grpc-rust example](https://github.com/stepancheg/grpc-rust/tree/master/grpc-compiler)