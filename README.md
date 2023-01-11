# Go cross-compile example

## Linux
```
GOOS=linux GOARCH=amd64 go build -o bin/hello-linux-amd64 hello.go
```

## Windows
```
GOOS=windows GOARCH=amd64 go build -o bin/hello-windows-amd64.exe hello.go
```

## Mac
### Intel-based Macs
```
GOOS=darwin GOARCH=amd64 go build -o bin/hello-darwin-amd64 hello.go
```

### M1-based Macs
```
GOOS=darwin GOARCH=arm64 go build -o bin/hello-darwin-arm64 hello.go
```
