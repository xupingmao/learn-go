package main


import (
	"fmt"
)


func panicAndRecover(arr []int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r);
		}
		ret = -1;
	}()

	var b int = arr[10];
	fmt.Printf("b = %v\n", b);
	return 0;
}


func main() {
	var arr = [...]int{1,2,3}

	var result = panicAndRecover(arr[:]);

	fmt.Printf("finished, result:%v\n", result);
}