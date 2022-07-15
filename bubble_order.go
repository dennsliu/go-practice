package main

import "fmt"

//冒泡排序
func main() {
	slice := []int{10, 16, 12, 8, 6}
	BubblingSort(slice)
	fmt.Println(slice)
}

func BubblingSort(arr []int) {
	var a int
	flag := false //标志位
	count := 1    //排序次数
	for _, _ = range arr {
		for ii, _ := range arr {
			if ii >= 1 && arr[ii] > arr[ii-1] {
				a = arr[ii-1]
				arr[ii-1] = arr[ii]
				arr[ii] = a
				flag = true
			}
			count++
			//fmt.Println(arr)
		}
		if flag == false {
			break
		}
		//fmt.Println(arr)
	}
	fmt.Printf("计算次数:%d", count)
}
