# runeweb
A Web service for finding Unicode characters by name

## Building

To compile `runeweb`, you need the `go-bindata` tool to bundle the Unicode database in the executable. To compile `runes`:

### 1 Get `go-bindata`

```
$ go get -u github.com/jteeuwen/go-bindata/...
```

### 2 Use it to generate the data file

This produces a `bindata.go` source file:

```
$ go-bindata data/
$ ls -la bindata.go
-rw-r--r--  1 lramalho  staff  3123912 Jan 29 18:09 bindata.go
```

### 3 Build and enjoy!

This compiles the `runes` self-contained executable:

```
$ go build
$ ls -l runes
-rwxr-xr-x  1 lramalho  staff  3205232 Jan 29 18:17 runes
$ ./runes party
U+1F389	🎉	PARTY POPPER
1 character found
```
