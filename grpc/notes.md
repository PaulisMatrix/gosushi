1.  Comman to generate grpc_pb.go and pb.go: `protoc --go_out=./grpc/hello_world --go_opt=paths=source_relative --go-grpc_out=./grpc/hello_world --go-grpc_opt=paths=source_relative ./grpc/hello_world/hello_world.proto `

2.  Run `go mod tidy` to fix the broken imports.

3.  Running protoc command will generate two files, `hello_world_grpc.pb.go` and `hello_world.pb.go`
    1. `hello_world.pb.go`, which contains all the protocol buffer code to populate, serialize, and retrieve request and response message types.<br>

    2. `hello_world_grpc.pb.go` which contains:
        a. An interface type (or stub) for clients to call with the methods defined in the `Greeter` service.<br>
        b. An interface type for servers to implement, also with the methods defined in the `Greeter` service.
