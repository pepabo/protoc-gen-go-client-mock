# protoc-gen-go-client-mock

`protoc-gen-go-client-mock` is a protoc plugin that generates a mock of [client that bundles gRPC Clients for multiple services](https://github.com/pepabo/protoc-gen-go-client).

## Example

Generate mock client ([client_mock.go](example/gen/go/myapp/client_mock.go)) that bundles gRPC clients generated by proto files in [the directory](example/proto/myapp).

## Requirements

- `protoc-gen-go`
- `protoc-gen-go-grpc`
- [`protoc-gen-go-grpc-mock`](https://github.com/sorcererxw/protoc-gen-go-grpc-mock)
- [`protoc-gen-go-client`](https://github.com/pepabo/protoc-gen-go-client)

## Options

| Option | Type | Description |
| --- | --- | --- |
| `package` | string | Specify package name (ex `--go-client_opt=package=xxxx` ) |
| `same_package` | bool | Make the package the same as the package generated by proto-gen-go (ex `--go-client_opt=same_package` ) |
| `with_close` | bool | Add `Close()` method for closing *grpc.ClientConn from client |
