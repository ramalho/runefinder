# runefinder

A tool for finding Unicode characters by name. Command-line and Web interfaces.

## Building

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


## Optional  Web interface

The `runeweb` command starts a local HTTP server on port 8000 offering a simple Web
interface for searching. This is the best way to use `runefinder` on Windows until
Microsoft improves the Unicode coverage of the fonts used in in cmd.exe or PowerShell.

Run the server:

```
$ cd cmd/runeweb/
$ go run runeweb.go
Serving on: localhost:8000
```

Open `http://localhost:8000` on your browser and search:

![Runefinder HTTP interface](https://github.com/standupdev/runefinder/raw/master/img/runefinder-runeweb.png "Runefinder HTTP interface")


## Web interface on Google App Engine

The `appengine/` directory has the `main.go` and configuration files for running Runefinder on the Google App Engine platform.

Link to running app: [runefinder2018.appspot.com](https://runefinder2018.appspot.com/)
