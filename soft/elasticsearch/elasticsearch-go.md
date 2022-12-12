# go-elasticsearch

## 示例
 
### 建立默认连接

```go
package main

import (
  "log"

  "github.com/elastic/go-elasticsearch/v7"
)

func main() {
  es, _ := elasticsearch.NewDefaultClient()
  log.Println(es.Info())
}
```

### 建立自定义连接

```go
package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	log.SetFlags(0)
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(es.Info())
		// => dial tcp: i/o timeout
	}
}

```