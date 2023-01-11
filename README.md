# Go cross-compile example

## Linux (64-bit)
```
GOOS=linux GOARCH=amd64 go build -o bin/hello-linux-amd64 hello.go
```

## Windows (64-bit)
```
GOOS=windows GOARCH=amd64 go build -o bin/hello-windows-amd64.exe hello.go
```

## Mac
### Intel-based Macs (64-bit)
```
GOOS=darwin GOARCH=amd64 go build -o bin/hello-darwin-amd64 hello.go
```

### M1-based Macs (64-bit)
```
GOOS=darwin GOARCH=arm64 go build -o bin/hello-darwin-arm64 hello.go
```

## Tips

- List all available architectures
  ```
  go tool dist list
  ```