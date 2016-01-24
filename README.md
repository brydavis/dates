# Dates

Quick date formatting based on [SQL datetime standards](http://www.w3schools.com/sql/func_convert.asp).

```go

package main

import (
	"fmt"
	d "github.com/brydavis/dates"
)

func main() {
	t1 := "7/21/1980 20:07:34.830"
	t2 := d.Convert(t1, "1/2/2006 15:04:05.000", 110)
	fmt.Println(t2) // 1980-07-21
}

```