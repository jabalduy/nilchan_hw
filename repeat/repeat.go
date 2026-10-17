package repeat

import "fmt"

func Repeat(arg string) {
	for range 5 {
		fmt.Println(arg)
	}
}
