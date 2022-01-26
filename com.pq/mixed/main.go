package main

import (
	"fmt"
	"math/rand"
)

func main() {
	testVarScope()
	testFuncVar()
}

/**
 */
func testFuncVar() {

}

/**
testVarScope 变量作用域
特殊的，在if/for中，可以声明与外部相同名称的局部变量，此时在作用域中将无法使用外部变量
*/
func testVarScope() {
	url := ""
	i := rand.Int()
	if i%2 == 0 {
		url, _ := test()
		fmt.Printf("url:%s pointer:%v\n", url, &url)
	}
	fmt.Printf("i:%d url:%s pointer:%v\n", i, url, &url)
}

func test() (string, error) {
	return "test", nil
}
