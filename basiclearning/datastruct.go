package basiclearning

import "fmt"

func ArrSliceExample() {
	var arr [5]int
	var arr2 [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	fmt.Println(arr2)

	slice1 := make([]int, 0)
	slice2 := make([]int, 3, 5)
	fmt.Println(len(slice2), cap(slice1))
	slice1 = append(slice1, 1, 2, 3, 4)
	sub1 := slice2[3:]
	//sub2... 是切片解构的写法，将切片解构为 N 个独立的元素。
	combined := append(sub1, slice1...)
	fmt.Print(combined)

}

func MapExample() {
	m1 := make(map[string]int)
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	m1["Tom"] = 18
	println(m1)
	println(m2)
}
