# Go gRPC server sample

以前までは以下のコマンド。
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    article/pb/article.proto
```

protocプラグインの変更により以下のコマンドでオプションを追加して実行する必要がある。
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --go-grpc_opt require_unimplemented_servers=false \
    article/pb/article.proto
```
参考資料: https://qiita.com/kohey_eng/items/5faaa82fda9fc15fa89d

## GraphQL

- https://zenn.dev/k88t76/articles/87bd2081ebcf97