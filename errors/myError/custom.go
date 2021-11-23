package main

import (
	"fmt"
	"math"
)

// error 错误类型是一个接口，只要实现了Error()这个方法，那么它就实现了这个接口
// error 不能直接用来比较

// 自定义错误
type radiusErr struct {
	radius float64
	err    string
}

func (e *radiusErr) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

// 计算园的面积
func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, &radiusErr{radius, "参数不能小于0"}
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
		
		if err, ok := err.(*radiusErr); ok {
			fmt.Println("捕捉自定义错误", err)
			return
		}
		fmt.Println("err: ", err)
        return
    }
    fmt.Printf("Area of circle %0.2f", area)
}