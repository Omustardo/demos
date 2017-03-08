Demo of basic gRPC usage. Loosely based on the hello world quickstart at http://www.grpc.io/docs/quickstart/go.html

Example usage:
```
go get github.com/omustardo/demos

# Open two terminals and go to this directory in each of them:
cd $GOPATH/src/github.com/omustardo/demos/grpc-calculator

# In one, run the server:
go run server/server.go

# In the other, run the client:
go run client/client.go
```
If it works, it will send an rpc from the client to the server, and get a response back.
The client should print out:
```
Results for input x:1 y:2
 Sum: 3
 Difference: -1
 Product: 2
 Quotient: 0.5
```

If you want to generate the gRPC service from the proto file, you need to have the proto compiler and the gRPC 
plugin installed. Follow my instructions for protoc and gRPC in my [setup docs](https://github.com/Omustardo/docs/blob/master/workspace/setup%20workspace.md)
if you don't have them installed yet.
The command to generate a `.pb.go` file from a `.proto` file that includes rpc services is:
```
protoc calc.proto --go_out=plugins=grpc:.
```
