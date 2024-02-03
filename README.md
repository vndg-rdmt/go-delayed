# go-delayed

### Install

```sh
go get -u github.com/vndg-rdmt/go-delayed
```

### Import

```go
import (
    delayed "github.com/vndg-rdmt/go-delayed"
)
```

### Usage

```go
d := delayed.New(delayed.Config{
    Epoch: time.Now(),
})

intervalTimer := d.SetInterval(time.Second * 5, func() {
    fmt.Println("This message printed each 5 seconds")
})

timeoutTimer := d.SetTimeout(time.Second * 5, func() {
    fmt.Println("This message printed after 5 seconds")
})
```