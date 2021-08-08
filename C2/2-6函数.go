package main

import "fmt"

func main() {
	if res, err := eval(1, 2, "*"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(res)
	}

	res2 := apply(div, 10, 2)
	fmt.Println(res2)

	res3 := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(res3)
}

//条件语句编写计算器
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func div(a, b int) int {
	return a / b
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}
