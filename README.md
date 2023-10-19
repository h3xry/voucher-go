## Truewallet Voucher Golang

### How to install
```go
go get github.com/h3xry/voucher-go
```
create file main.go
```go
package main

import (
	"fmt"

	"github.com/h3xry/voucher-go"
)

func main() {
	ret, err := voucher.Redeem("0999999999", "https://gift.truemoney.com/campaign/?v=1982376419827369830569i0c9uhKJHkl203")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}

```
run command
```go
go run main.go
```

**Enjoy your service**

## 🧧 Donate me

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/h3xry)