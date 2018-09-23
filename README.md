# intds
`intds` stands for integer data server. It is a client/server tool for sending arrays of integers to a server and receiving arrays of integers in return. That's seriously all it does.

## usage
The `intds` CLI starts a client or server, both of which live on your machine as a gRPC server. The client may only be used to send data, while the server may be used to both listen for data, and send responses.

### server
The `intds` server listens for a connection and, upon establishing that connection, will stream data from the client. 
```
$ intds server --port 8080
```
The gRPC calls can be generated so that once a server is running, it may be used to both send and receive data.

### client
The `intds` client can only send data to the server.
```
$ intds client --host <server IP/hostname> --port <server port>
```
The client gRPC calls then may be used to send data to the server.
