package channel_test

import "fmt"

func ExamplePoolFunc() {
	fn := chanel.PoolFunc(2)

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
