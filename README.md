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

Currently, the only buildable executable is the `runes` command, located in the `cmd/runes` directory.

To compile the `runes` command:

```
$ cd cmd/runes/
$ go build
$ ls -lh runes
-rwxrwxr-x 1 luciano luciano 2,6M Mar  7 00:16 runes
```

To use `runes`, provide one or more words to search:

```
$ ./runes party
U+1F389	🎉	PARTY POPPER
1 character found

$ ./runes cat eyes
U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
U+1F63D	😽	KISSING CAT FACE WITH CLOSED EYES
3 characters found
```
