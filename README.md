<p align="center">
<img src="./assets/photo/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# queue 
queue is a lightweight package for having queue and deque in Go.
it performs better than the built-in [container/list](https://golang.org/pkg/container/list/) package for any operation.

# Documentation

## Install

```bash
go get github.com/erfanmomeniii/queue
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/queue"
``` 

## Quick Start
The following example illustrates how to use this package for creating an instance and performing any operation with it:
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/queue"
)

func main() {
	q := queue.New()
	q.PushFront(1)
	q.PushFront("hi")
	q.PushBack(3.14)
	// q => [ "hi", 1, 3.14]

	fmt.Println(q.Front())
	// hi
	fmt.Println(q.Back())
	// 3.14

	q.PopBack()
	q.PopFront()
	fmt.Println(q.Front())
	// 1
}

```
## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
