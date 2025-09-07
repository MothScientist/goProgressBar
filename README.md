# goProgressBar

Very simple and easy to use progress bar for console

## Examples

```go
package main

import (
    "" // TODO
    "fmt"
    "time"
)

func main() {
    bar := progressbar.GetNewProgressBar()
    bar.SetColors([2]string{progressbar.ColorWhite, progressbar.ColorPurple})
    bar.SetBarLen(100)
    for i := uint8(0); i <= 100; i++ {
        fmt.Println(bar.Update(i))
        time.Sleep(300 * time.Millisecond)
    }
}
```
