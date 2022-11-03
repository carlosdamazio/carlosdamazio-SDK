## LOTR SDK
This is a Lord Of The Rings API SDK written in `go`. It fully supports the original
LOTR API: [The One API](https://the-one-api.dev/).

## Usage
The only thing you'll need is the API key, which you can grab one from the original's
API website.

```bash
export GO_LOTR_API=<your api key>

go get github.com/carlosdamazio/carlosdamazio-SDK
```

```go
package main

import (
	"fmt"

	"github.com/carlosdamazio/lotrsdk/pkg/api"
)

func main() {
	client := api.New()

	books, err := client.Book.List()
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Printf("ID: %s, Name: %s\n", book.ID, book.Name)
	}
}
```
