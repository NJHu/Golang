## Golang基础

- 变量定义

```go
import (
	"fmt"
	"math"
	"reflect"
)

//math
func testMath() {
	var value float64 = 1.2
	var num float64 = math.Pow(value, 100)
	fmt.Println(num)
}
```

- 自动推导

```go
import (
	"fmt"
	"math"
	"reflect"
)
//自动推导
func autoType()  {
	a := 10
	b := 10.12
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
}

```

- 输出格式化

```go
a := 10
b := 10.16
fmt.Printf("%d, %.1f", a, b)
```


- 输出类型

```go
func main() {
	a := 10
	b := 10.1
	c := "123"
	d := true

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}

int
float64
string
bool
```

- 输入

```go
func main() {
	var weight float64
	fmt.Println("请输入单价")
	// 输入
	fmt.Scan(&weight)

	price := 3.2
	sum := weight * price

	fmt.Println(sum)
	fmt.Printf("%.2f", sum)
}
```

- 字符

```go
func main0201() {
	//byte字符类型 同时也是uint8的别名
	var a byte = 'a'

	//所有的字符都对应ASCII中的整型数据
	//'0'对应的48  'A'对应的65 'a' 对应的97
	//fmt.Println(a)

	//%c是一个占位符 表示打印输出一个字符
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", 97)

	fmt.Printf("%T\n", a)
	var b byte = '0' //字符0  对应的ASCII值为为48

	fmt.Printf("%c\n", 48)
	fmt.Printf("%c\n", b)
}
func main0202() {

	var a byte = 'a'
	//将小写字母转成大写字母输出
	fmt.Printf("%c", a-32)
}

func main() {
	//转义字符 \n 换行

	//var a byte = '\n'
	//\0  对应的ASCII 值为0 用于字符串的结束标志
	//\t 对应的ASCII 值为9 水平制表符 一次跳八个空格
	var a byte ='\t'
	//fmt.Println(a)
	fmt.Printf("%c",a)
}

```



















