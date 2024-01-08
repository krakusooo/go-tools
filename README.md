# Go Tools — CLI Utilities

A collection of command-line tools written in Go.

## Tools

### portscanner
Fast concurrent TCP port scanner.
```bash
go run portscanner/main.go -host 192.168.1.1 -ports 1-1024
```

### logparser
Parses and filters system/web server logs.
```bash
go run logparser/main.go -file /var/log/nginx/access.log -filter "404"
```

### hashcheck
File integrity checker using MD5/SHA256.
```bash
go run hashcheck/main.go -file document.pdf -algo sha256
```

## Build
```bash
go build ./...
```

## Requirements
Go 1.21+
