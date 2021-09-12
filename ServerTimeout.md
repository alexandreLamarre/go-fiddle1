## HTTP Server Timeouts

- Setting timeouts in server is important to conserve system resources and to protect from DDOS attack
- File descriptors are limited
- Malicious user can open many client connections, consuming all file descriptors
- Server will not be able to accept any new connection


### net/http Timeouts

- There are four main timeouts exposed in http.server
  - Read Timeout
  - Read Header Timeout
  - Write Timeout
  - Idle Timeout

### Set Timeouts by explicitly using a Server

```
srv := &http.Server{
    ReadTimeout: 1 * time.Second,
    ReadHeaderTimeout: 1 * time.Second,
    WriteTimeout: 1 * time.Second,
    IdleTimeout: 30 * time.Second,
    Handler: serveMux,
}
```

- Set Connection timeouts when dealing with untrusted clients and networks
- Protect Server from clients which are slow to read and write

### HTTP Handler Functions

- Connection timeouts apply at network connection level
- HTTP Handler Functions are unaware of these timeouts, they run to completion, consuming resources


#### Context Timeouts and Cancellation
- Use Context timeouts and cancellation to propagate the cancellation signal down the call graph
- The Request Type already has a context attached to it

```
ctx := req.Context()
```

- Server cancels thi context when,
  - Client closes the connection
  - Timeout
  - ServeHTTP method returns