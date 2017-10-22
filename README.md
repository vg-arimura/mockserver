# mockserver

Tiny mock server for you

## How to use
```
$ echo "hoge" > data/hoge
$ go run main.go -data=data -port=8000
$ curl http://localhost:8000/hoge
```
