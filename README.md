# runefinder

A tool for finding Unicode characters by name. Command-line and Web interfaces.

## Building

To compile `runefinder`, you need the `go-bindata` tool to bundle the Unicode database in the executable. To compile `runes`:

### 1 Get `go-bindata`

```
$ go get -u github.com/jteeuwen/go-bindata/...
```

### 2 Use it to generate the data file

This step produces a `bindata.go` source file:

```
$ go-bindata -pkg runefinder data/
$ ls -lh bindata.go 
-rw-rw-r-- 1 luciano luciano 1014K Mar  7 00:10 bindata.go
```

### 3 Build and enjoy!

To compile the `runes` command, in the `cmd/runes` directory:

```
$ cd cmd/runes/
$ go build
$ ls -lh runes
-rwxrwxr-x 1 luciano luciano 2,6M Mar  7 00:11 runes
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
