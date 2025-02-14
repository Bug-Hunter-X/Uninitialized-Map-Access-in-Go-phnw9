```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, it will print 0
    fmt.Println(m["b"]) // This will not panic, it will print 0
    var n *map[string]int
    fmt.Println((*n)["a"])
}
```