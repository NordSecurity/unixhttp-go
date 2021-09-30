# unixhttp
## Description
Simple library with helpers for making HTTP clients/servers that communicate over Unix domain socket.

## Usage

### Echo Listener

Substitute `Listener` for the default `http` one to listen on a unix socket:

```go
unixSocket := "/tmp/test.sock"

listener, err := unixhttp.NewListener(unixSocket)
if err != nil {
    panic(err)
}
defer listener.Close()

api := echo.New()
api.Listener = listener
api.Start("")
```

### HTTP Client

Use `Client` to make HTTP calls over a unix socket.

```go
unixSocket := "/tmp/test.sock"

client := unixhttp.NewClient(unixSocket)
response, err = client.Get("https://127.0.0.1/test")
```
