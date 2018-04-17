# imageconverter
image converter in Go

## Usage
basic usage
```
$ go run main.go -outputformat <outputformat> ./<yourimagefolder>
```
this tool can convert to png, jpg or ascii file
```
$ # convert to png
$ go run main.go -outputformat png ./<yourimagefolder>
```

```
$ # convert to ascii
$ go run main.go -outputformat ascii ./<yourimagefolder>
```