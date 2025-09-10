# goProgressBar

Very simple and easy to use progress bar for console

## Examples

```
go get github.com/MothScientist/goProgressBar@latest
```

```go
package main

import (
	"fmt"
	"time"

	progressbar "github.com/MothScientist/goProgressBar"
)

func main() {
	bar := progressbar.GetNewProgressBar()

	bar.SetSpinner(progressbar.Spinners[1])
	// There is no point in setting the color white in the console,
	// you can save bytes by setting an empty string instead
	bar.SetColors([2]string{"", progressbar.ColorPurple})
	bar.SetBarLen(100)

	for i := 30; i <= 100; i++ {
		fmt.Println(bar.Update(i))
		time.Sleep(300 * time.Millisecond)
	}
}
```

![example1.png](images/example1.png)
![example2.png](images/example2.png)
![example3.png](images/example3.png)
![example4.png](images/example4.png)