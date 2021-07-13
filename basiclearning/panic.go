package basiclearning

import "fmt"

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("nono")
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	fmt.Println(get(5))
	fmt.Println("finished")
}
