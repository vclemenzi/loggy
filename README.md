![](https://i.imgur.com/T0L9C5v.png)

# Loggy
Better `fmt.Println` for go. Inspired by [consola](https://github.com/unjs/consola)

```go
package main

import "github.com/vclemenzi/loggy"

func main() {
  loggy.Info("Hello, world!", "This is a log message.")
  loggy.Error("Hello, world!", "This is a log message.")
  loggy.Warn("Hello, world!", "This is a log message.")
  loggy.Success("Hello, world!", "This is a log message.")
  loggy.Box("Hello, world! This is a log message.")
}
```
