# gcm
GCM client for golang

## How to use

Gets library

```
$ go get github.com/mokelab-go/gcm
```

write!

```go
package main

import (
  "fmt"
  gcm "github.com/mokelab-go/gcm/impl"
)

const (
  api_key = "INPUT YOUR API_KEY"
  reg_id = "INPUT YOUR REG_ID"
)

func main() {
  client := gcm.NewClient()
  data := map[string]interface{}{}
  _, err := client.Send(api_key, data, reg_id)
  if err != nil {
    fmt.Errorf("Failed to send GCM message : %s", err)
  }
}
```

