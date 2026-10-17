package repeat

import "fmt"

func Repeat(count int, arg any) {
	for range count {
		fmt.Println(arg)
	}
}
