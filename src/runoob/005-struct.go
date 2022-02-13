package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func (c *Circle) Description() string {
	return fmt.Sprintf("{Center=%v, Radius=%v}", c.Point, c.Radius)
}

func main() {
	// 匿名成员无法通过属性来初始化
	// unknown field 'X' in struct literal of type Circle
	// unknown field 'Y' in struct literal of type Circle
	// var c = Circle{X:0,Y:0,Radius:3}

	// 必须这样声明
	var c = Circle{Point{0,0}, 3}
	var c2 = Circle{Point{0,0},3}
	var c3 = Circle{Point{1,1},3}
	fmt.Printf("c=%v\n", c) 
	fmt.Printf("c.description=%v\n", c.Description())

	// 如果结构体的每个属性都能进行比较，那它就可以被进行比较
	fmt.Printf("c==c2:%v\n", c == c2)
	fmt.Printf("c==c3:%v\n", c == c3)
}