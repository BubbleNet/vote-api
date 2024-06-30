# vote-api

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