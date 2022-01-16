// 温度转换的包
package tempconv

import "fmt"

// type语法
// type name underlying-type

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToString() string {
	return fmt.Sprintf("%gC", c)
}

func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit((c*9/5+32))
}

func (f Fahrenheit) ToString() string {
	return fmt.Sprintf("%gF", f)
}

func (f Fahrenheit) FToC() Celsius {
	return Celsius((f-32)*5/9)
}


// Go语言约定的包初始化函数
func init() {
	fmt.Printf("tempconv init\n")
}

// 可以有任意多个init函数
// 按照声明的顺序自动执行
func init() {
	fmt.Printf("tempconv init2\n")
}
