# envflag

Go package that extends `flag` allowing flags to be overriden with environment variables.

Example usage:

```go
package main

import (
	"flag"
	"log"

	"github.com/igorsobreira/envflag"
)

func main() {
	var url, secret string

	flag.StringVar(&url, "url", "http://example.com", "API URL")
	envflag.StringVar(&secret, "secret", "SECRET", "", "API Secret")

	flag.Parse()
	flag.VisitAll(envflag.Visit)  // override flag values with env vars

	log.Printf("url = %q", url)
	log.Printf("secret = %q", secret)
}
```
