https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/codegenerate

```bash
➜  api git:(main) ✗ protoc --go_out=../pkg/grpc --go_opt=paths=source_relative --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative hello.proto
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
```

getではなくinstallした。

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

あとはパスに `export PATH="$PATH:$(go env GOPATH)/bin"` を追加した。