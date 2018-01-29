# runes

A command-line utility to find Unicode characters by name

## Usage

To find Unicode characters, run `runes` with words as arguments:

```
$ ./runes cat face
U+1F431	🐱	CAT FACE
U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
U+1F639	😹	CAT FACE WITH TEARS OF JOY
U+1F63A	😺	SMILING CAT FACE WITH OPEN MOUTH
U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
U+1F63C	😼	CAT FACE WITH WRY SMILE
U+1F63D	😽	KISSING CAT FACE WITH CLOSED EYES
U+1F63E	😾	POUTING CAT FACE
U+1F63F	😿	CRYING CAT FACE
U+1F640	🙀	WEARY CAT FACE
10 characters found
```

Use more words to narrow the results:

```
$ ./runes cat face eyes
U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
U+1F63D	😽	KISSING CAT FACE WITH CLOSED EYES
3 characters found
```

## Building

To compile `runes`, you need the `go-bindata` tool to bundle the Unicode database in the executable. To compile `runes`:

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

