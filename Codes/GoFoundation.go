//文件所属的包 在go语言中 主函数所在的包一定是main
package main

import "fmt"

func main() {

	//fmt.Printf("%T", sumArr);
	fmt.Println(demo1())
}

func sumArr(arrNums...int) (sum int) {
	fmt.Printf("%T", arrNums)

	for _, v := range arrNums {
		println(v)
	}
	sum = 10
	return
}

func demo1() int {
	return 21
}



