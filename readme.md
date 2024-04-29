## base on logrus, realizing customize log

### example

```go
package main

import "github.com/Kress4s/kslog"

func main() {
	// log setting
	kl := kslog.New("your app name")
	kl.SetFormatter(&kslog.PrettyJSONFormatter)
	kl.SetLevel(kslog.TranceLevel)

	// log
	kl.Debugf("test log, app: %s", "test")
	kl.Traceln("test")
	// ...
}
```