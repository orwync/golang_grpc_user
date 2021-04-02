#gRPC using golang to retrive users

dependecies:

```bash
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf
```

run tests:

```bash
$ go test ./user/user_server/
```

run to start:

```bash
$ go run user/user_server/server.go user/user_server/helper.go user/user_server/db.go user/user_server/models.go
```

The server will start on port 50051

To connect to the server user:
Run:

```bash
$ go run user/user_client/client.go
```

OR

Evans:
[Evans installation](https://github.com/ktr0731/evans#installation)

        Run: $ evans -p 50051 -r

        In Evans CLI:
            To get user by User Id: call UserById
                Enter the Id

            To get users by list of Ids: call UserByListId
                Enter the User ID one by one
                After entering all user id: Ctrl+D
