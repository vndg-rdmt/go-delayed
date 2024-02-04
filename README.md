# go-delayed [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

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

[ci-img]: https://github.com/vndg-rdmt/go-delayed/actions/workflows/ci.yml/badge.svg

[ci]: https://github.com/vndg-rdmt/go-delayed/actions/workflows/ci.yml

[cov-img]: https://codecov.io/gh/vndg-rdmt/go-delayed/branch/main/graph/badge.svg?token=h2KPSXfm5E

[cov]: https://codecov.io/gh/vndg-rdmt/go-delayed