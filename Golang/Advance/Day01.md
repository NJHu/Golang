## 指针

```go
package main
import "fmt"

func main()  {
    var a int = 10
    var p *int  = &a
    fmt.Println(&a, a, p, &p, *p)
}

/*log
&a = 0xc000092000
a = 10
p = 0xc000092000
&p = 0xc00008c008
*p = 10
*/
```

## 栈帧
> 用来给函数运行提供内存空间, 取内存于stack

- 1, 局部变量
- 2, 形参
- 3, 内存字段描述值
  - 栈基指针, 栈顶指针

## 空指针和野指针
> C语言

## new
> 在heap堆上申请空间









