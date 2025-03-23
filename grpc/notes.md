1.  Comman to generate grpc_pb.go and pb.go: `protoc --go_out=./grpc/hello_world --go_opt=paths=source_relative --go-grpc_out=./grpc/hello_world --go-grpc_opt=paths=source_relative ./grpc/hello_world/hello_world.proto `

2.  Run `go mod tidy` to fix the broken imports.

3.  Running protoc command will generate two files, `hello_world_grpc.pb.go` and `hello_world.pb.go`
    1. `hello_world.pb.go`, which contains all the protocol buffer code to populate, serialize, and retrieve request and response message types.<br>

    2. `hello_world_grpc.pb.go` which contains:
        a. An interface type (or stub) for clients to call with the methods defined in the `Greeter` service.<br>
        b. An interface type for servers to implement, also with the methods defined in the `Greeter` service.


## Some helpful tips while structing protos

* Adhere to these best practices to avoid any confusion : https://buf.build/docs/reference/protobuf-files-and-packages/

* The name `package <name>` should ideally follow the directory structure of your proto file. For ex:

      proto
          -myorg
              -chat
                  -v1
                    -feedback.proto
  Then the package definition should be `package proto.myorg.chat.v1`

* `option go_package` signifies the import path of the req/res files in case of golang. This is directory in which protoc will place the `_pb.go` 
   files.  For ex: `option go_package = "go/chat/v1;chatv1"`, assuming you have created go/chat/v1 directory at the root. 

   Then the overall import path becomes : `chatv1 github.com/paulismatrix/practice/go/chat/v1` with
   package name as `chatv1`

      go
        -chat
           -v1
             -feedback_pb.go
      proto  

* If you don't specify directory in `--go_out=.` flag then protoc will use `go_package` path to place the generated files.

  protoc cmd becomes : `protoc --proto_path=. --go_out=. proto/myorg/chat/v1*.proto`

* Generally, it's a good convetion to structure the protos directory path based on different versions like `/v1`, `/v2`, etc to maintain all the 
  versions together and to support backward compatibility. 
