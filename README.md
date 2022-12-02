# html-link-parser-go

A package that parse html link to Go struct

## Install

```
go get -u github.com/tanerijun/html-link-parser-go/parser
```

## Examples

```go
func main() {
    f, err := os.Open("example.html")
	if err != nil {
		panic(err)
	}

	links := parser.Parse(f)
}
```
