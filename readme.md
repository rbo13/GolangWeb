# Golang Web

Golang Web is my sample web app written in Go. This is a server side rendered golang application.

# Learnings:
  - Creating smart templates for efficient view rendering.
  - Working with HTTP Requests.
  - Creating custom middleware in golang.
  - Interacting with a Database (MySQL)
  - Uses HTTP/2
  - HTTP/2 Server.Push
  
### Tech
  - Golang
  - MySQL
  - Unit Testing (Go Built-in test)
  - Performance Profiling:
    * go tool pprof http://localhost:8000/debug/pprof/heap //memory
    * go tool pprof http://localhost:8000/debug/pprof/profile //cpu
    * go tool pprof http://localhost:8000/debug/pprof/block //goroutines
    * go tool pprof http://localhost:8000/debug/pprof/trace?seconds=5 //trace
    * go tool pprof http://localhost:8000/debug/pprof //Web Interface

### Tests
- Please see each directory prefixed with `_test`.
    * Running Tests:
        * `go test github.com/lss/controller` / (where the test file is located) 

### Installation

Assuming you have the right path set in your golang workspace.

Clone this repo to your workspace.

```sh
$ cd src/github.com/lss
$ npm install (yes, some of my assets are built with node)
$ go run main.go
```

Alternative (using https://github.com/pilu/fresh)

```sh
$ cd src/github.com/lss
$ npm install
$ fresh
```

### Todos

 - Write MORE Tests

License
----

MIT
