# gRPC Docker TODO

## Commands
```
protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt require_unimplemented_servers=false api/todopb/todo.proto
```

**without grpc**
```
protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_opt require_unimplemented_servers=false api/todopb/todo.proto
```

## References
- https://selfnote.work/20200810/programming/todoapp-with-go-grpc-docker/
- 