<h1>变量作用域</h1>

根据变量定义位置的不同，可以分为以下三个类型：

函数内定义的变量称为局部变量

函数外定义的变量称为全局变量

函数定义中的变量称为形式参数

<h2>局部变量</h2>

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，函数的参数和返回值变量都属于局部变量。

局部变量不是一直存在的，它只在定义它的函数被调用后存在，函数调用结束后这个局部变量就会被销毁。

【示例】下面的 main() 函数中使用到了局部变量 a、b、c。
package main
import (
    "fmt"
)
func main() {
    //声明局部变量 a 和 b 并赋值
    var a int = 3
    var b int = 4
    //声明局部变量 c 并计算 a 和 b 的和
    c := a + b
    fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
运行结果如下所示：
a = 3, b = 4, c = 7
<h3><color="red">**特别的**</h3>
<h2>全局变量</h2>

在函数体外声明的变量称之为全局变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，当然，不包含这个全局变量的源文件需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。

全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。

【示例】下面代码中，第 6 行定义了全局变量 c。
package main
import "fmt"
//声明全局变量
var c int
func main() {
    //声明局部变量
    var a, b int
    //初始化参数
    a = 3
    b = 4
    c = a + b
    fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
运行结果如下所示：
a = 3, b = 4, c = 7
Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
package main
import "fmt"
//声明全局变量
var a float32 = 3.14
func main() {
    //声明局部变量
    var a int = 3
    fmt.Printf("a = %d\n", a)
}
运行结果如下所示：
a = 3
<h2>形式参数</h2>

在定义函数时函数名后面括号中的变量叫做形式参数（简称形参）。形式参数只在函数调用时才会生效，函数调用结束后就会被销毁，在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值。

形式参数会作为函数的局部变量来使用。

【示例】下面代码中第 21 行定义了形式参数 a 和 b。
package main
import (
    "fmt"
)
//全局变量 a
var a int = 13
func main() {
    //局部变量 a 和 b
    var a int = 3
    var b int = 4
    fmt.Printf("main() 函数中 a = %d\n", a)
    fmt.Printf("main() 函数中 b = %d\n", b)
    c := sum(a, b)
    fmt.Printf("main() 函数中 c = %d\n", c)
}
func sum(a, b int) int {
    fmt.Printf("sum() 函数中 a = %d\n", a)
    fmt.Printf("sum() 函数中 b = %d\n", b)
    num := a + b
    return num
}
运行结果如下所示：
main() 函数中 a = 3
main() 函数中 b = 4
sum() 函数中 a = 3
sum() 函数中 b = 4
main() 函数中 c = 7
在for/if结构中