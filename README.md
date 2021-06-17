# Exemplo gRPC
## Implementação grpc utilizando GO

### Gerando o arquivo .proto
> protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

### Rodando o Server
> go run .\cmd\server\server.go

### Rodando o Client
> go run .\cmd\client\client.go
