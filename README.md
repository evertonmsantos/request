## Módulo de Requisições HTTPs

##### Como instalar

```
go get github.com/evertonmsantos/request
```

##### Código exemplo usando o método GET

```go
package main

import (
    "fmt"
    
    "github.com/evertonmsantos/request"
)

func main() {

    req, err := request.Get("https://httpbin.org/get", make(map[string]string), false)
    
    if err != nil {
        panic(err)
    }
    
    fmt.Println(req)

}
```

##### Retorno

```
&{200 OK 200 240 {
  "args": {},
  "headers": {
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-62fd1be9-69f4867c2900732c45c4421c"
  },
  "origin": "",
  "url": "https://httpbin.org/get"
}
 map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[*] Content-Length:[240] Content-Type:[application/json] Date:[Wed, 17 Aug 2022 16:48:41 GMT] Server:[gunicorn/19.9.0]]}
```
