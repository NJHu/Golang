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

- 字符串

```go
    var str1 string = "1第一句1"
	var str2 string = "2第二句2"
	fmt.Println(str1 + str2)
	fmt.Println(len(str1))
	fmt.Printf("%d", len(str2))
```

- 进制

```go
	var a int = 225
	fmt.Printf("二进制: %b\n", a)
	fmt.Printf("八进制: %o\n", a)
	fmt.Printf("十进制: %d\n", a)
	fmt.Printf("十六机制大: %X\n", a)
	fmt.Printf("十六进制小: %x\n", a)

	var b int = 010;
	var c int = 0xABC;
	fmt.Println(b)
	fmt.Println(c)

	//
	二进制: 11100001
    八进制: 341
    十进制: 225
    十六机制大: E1
    十六进制小: e1
    8
    2748
```

- 枚举

```go
	//在定义iota枚举时可以自定义赋值
	const (
		a = iota
		b
		c = 20
		d
		e
		f = iota
		g
	)
	//a=100//err
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	//
	0
    1
    20
    20
    20
    5
    6
```

- fallthrough

```go
//根据输入的年份月份 计算这个月有多少天

	var y int
	var m int
	fmt.Scan(&y, &m)

	//在switch语句中可以把相同的值放在一个case中
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(30)
	//fallthrough 在case中向下执行下一个case
	//case 1:
	//	fallthrough
	//case 3:
	//	fallthrough
	//case 5:
	//	fallthrough
	//case 7:
	//	fallthrough
	//case 8:
	//	fallthrough
	//case 10:
	//	fallthrough
	//case 12:
	//	fmt.Println(31)
	//
	//case 4:
	//	fallthrough
	//case 6:
	//	fallthrough
	//case 9:
	//	fallthrough
	//case 11:
	//	fmt.Println(30)
	case 2:
		//判断是否是闰年  能被4整除 但是 不能被100整除  或 能被400整除
		if y%4 == 0 && y%100 != 0 || y%400 == 0 {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}

	default:
		fmt.Println("月份输入错误")
	}
```

- `goto`and`FLAG`

```go
	fmt.Println("hello world1")
	fmt.Println("hello world2")
	//如果在代码中入到goto 会跳到所定义的标志位
	//可以在一个循环中跳到另外一个循环中  可以在一个函数中跳到另外一个函数中
	goto FLAG
	fmt.Println("hello world3")
	fmt.Println("hello world4")
	FLAG:
	fmt.Println("hello world5")
	fmt.Println("hello world6")
```

- 数组遍历

```go
func sumArr(arrNums...int) {
	for _, v := range arrNums {
		println(v)
	}
}
```

- 函数返回值

```go
//多个返回值
func test5(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
```

- 函数类型

```go
type FuncType func(int, int)
```

- 匿名单数

```go
var f FuncType
f = func(a int, b int) {
	fmt.Println(a + b)
}
f(a, b)
```






