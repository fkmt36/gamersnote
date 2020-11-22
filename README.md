# GamersNote

## go-swagger

https://github.com/go-swagger/go-swagger

以下のコマンドで定義ファイルからAPI生成

```
swagger generate server -A gamersnote-api -P handlers.TokenPayload -t ./api -f ./swagger/swagger.yml
```

以下のコマンドで定義ファイルからAPI Clientを生成

```
openapi-generator generate -g typescript-axios -i ./swagger/swagger.yml -o ./app/api-client
```
