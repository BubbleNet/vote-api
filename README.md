# vote-api

vote-api is a REST API used to access voter registration infornmation and election data.

## Install dependancies

```
go mod download
```

## Build

```
go build ./cmd/main.go
```

## Run

```
go run ./cmd/main.go
```
or
```
go run ./main
```
after building

## Unit Tests

```
go test ./...
```

## Integration Tests

(with the server running)

```
go test ./... --tags=integration
```

## VSCode Setup
go-delve is required for VSCode to run this service.
```
go install -v github.com/go-delve/delve/cmd/dlv@latest
```
