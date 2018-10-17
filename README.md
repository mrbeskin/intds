# intds
`intds` stands for integer data server. It is a client/server tool for sending arrays of integers to a server and receiving arrays of integers in return. That's seriously all it does.

## usage
The `intds` CLI starts a client or server, both of which live on your machine as a gRPC server. The client may only be used to send data, while the server may be used to both listen for data, and send responses.

### server
The `intds` server listens for a connection and, upon establishing that connection, will stream data from the client. 
```
$ intds server -p <port> -g <gRPC port>
```
The gRPC calls can be generated so that once a server is running, it may be used to both send and receive data.

The port for the `-p` flag is the port on which client connections are accepted and thus should correspond to the same flag on the corresponding client. 

The `-g` flag is the one that the gRPC client will use to connect to this service. This defaults to `50051`.

### client
The `intds` client can only send data to the server.
```
$ intds client -H <server IP/hostname> -p <server port> -g <gRPC port>
```
The client gRPC calls then may be used to send data to the server.

The port for the `-p` flag is the port on which the server is accepting client connections and thus should correspond to the same flag on the server. 

The `-g` flag is the one that the gRPC client will use to connect to this service. This defaults to `50051`.

### Generating gRPC 

Definitions are included for Go and Python . For clients in other languages, please refer to the documentation on the [grpc site](https://grpc.io) for more information.

For Python, the directions for creating a client are located [here](https://grpc.io/docs/tutorials/basic/python.html).

For C#, the directions for creating a client are located [here](https://developers.google.com/protocol-buffers/docs/reference/csharp-generated)
