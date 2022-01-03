# Domain Generation Algorithm
## Run
```
go run src/main.go

# Example
go run src/main.go -k 24MJSMB3FDDZOAZC -t 1641216911
```

## Build CLI
```
go build -buildmode=exe -o dist/dga src/main.go
```

## Build C Shared Libary
```
go build -buildmode=c-shared -o lib/libdga.so src/libDga/libDga.go
```
