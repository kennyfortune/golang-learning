package basiclearning

import (
	"fmt"
	"reflect"
)

/**
	内置类型

	空值：nil
	整型类型： int(取决于操作系统), int8, int16, int32, int64,
			  uint8, uint16, uint32, uint64, uintptr
	浮点数类型：float32, float64
	复数类型：complex64 complex128
	字节类型：byte (等价于uint8)
	字符串类型：string
	布尔值类型：bool，(true 或 false)

	byte // uint8 的别名
	rune // int32 的别名, 表示一个 Unicode 码点
**/

// 声明
var a int                           // 声明
var b int = 1                       // 声明并赋值
var c = 1                           // 声明并赋值，自动推断类型,
var python, java bool = true, false // 声明时同样类型的变量声明，可以省略前面的类型
var i complex64 = complex64(1 + 2i)

func declare() {
	me := true             // 只有一个变量时，声明并赋值
	d := 1                 // 函数外的每个语句都必须以关键字开始（`var`, `func` 等等）
	msg := "hello"         // 因此 `:=` 结构不能在函数外使用。
	me, you := false, true // 多个变量时，如果变量已声明则为赋值修改，未声明则是声明并赋值
	fmt.Println(me, d, msg, you)
}

// 常量
const Pi float32 = 3.14 //常量可以是字符、字符串、布尔值或数值, 量不能用 `:=` 语法声明。
const (
	// 常量可以用表达式声明，若没有指定类型，则它为 untyped 常量
	// untyped 的数值常量是高精度的，赋值到低精度的类型会编译错误
	Big   = 1 << 80
	Small = Big >> 99
	// testa int = 1208925819614629174706176
	// testb int = Big
	// 上面两行会提示overflow， 1208925819614629174706176，Big 在这都为 untyped数值常量
)

// 类型转换
func typecasting() {
	var i int = 42 // golang是强类型语言，使用T(v)进行类型转换
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Print(u)

	// string使用 UTF8 编码，英文每个字符占1byte，如果是中文，一般占3字节。
	// 转换成 []rune 类型后，字符串中的每个字符，无论占多少个字节都用 int32 来表示，因而可以正确处理中文。
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind())    // uint8
	fmt.Println(str1[2], string(str1[2]))          // 108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])        // 232 è
	fmt.Println("len(str2)：", len(str2))           // len(str2)： 8
	runeArr := []rune(str2)                        // 类型转换成数组
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
	fmt.Println("len(runeArr)：", len(runeArr))     // len(runeArr)： 4
}
