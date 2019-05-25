//文件所属的包 在go语言中 主函数所在的包一定是main
package main

import (
	"fmt"
)

func main() {

	var a byte = 'a'
	fmt.Printf("%T\n", a)

	fmt.Printf("%c\n", a)
	fmt.Printf("%d\n", a)
}



