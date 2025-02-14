```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, it will print 0
    fmt.Println(m["b"]) // This will not panic, it will print 0
    m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["a"]) // This will print 1
    //Safe way
    var n *map[string]int
    if n != nil {
        fmt.Println((*n)["a"])
    }
}
```