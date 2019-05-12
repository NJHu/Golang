## 指针

```go
package main
import "fmt"

func main()  {
    var a int = 10
    var p *int  = &a
    fmt.Println(&a, a, p, &p, *p)
}

// log
&a = 0xc000092000
a = 10
p = 0xc000092000
&p = 0xc00008c008
*p = 10
```


