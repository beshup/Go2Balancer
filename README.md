**Concurrent TCP Load Balancer**


## Startup
```
cd app
```
```
go run . --serveOn=3000 --portBackends=5000/8080 --urlBackends=golang.org/example.com --ipBackends=192.0.2.1
```

## HTTP2 Support/Switch Coming Soon

This load balancer was originally meant to interface HTTP2 clients with HTTP1 servers.
