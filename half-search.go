package main

import "fmt"

/*
二分查找函数，假设有序数组的顺序是从小到大
*/
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findValue int) {

	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("未找到")
		return
	}
	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findValue {
		fmt.Printf("小于中间 middle：%d, leftIndex： %d rightIndex：%d findValue：%d ", middle, leftIndex, middle-1, findValue)
		//要查找的数，范围应该在leftIndex与 middle-1 之间
		BinaryFind(arr, leftIndex, middle-1, findValue)

	} else if (*arr)[middle] < findValue {
		fmt.Printf("大于中间  middle：%d, leftIndex： %d rightIndex：%d findValue：%d ", middle, middle+1, rightIndex, findValue)
		//要查找的数，范围应该在middle+1与 rightIndex 之间
		BinaryFind(arr, middle+1, rightIndex, findValue)

	} else {
		fmt.Printf("找到了，下标为：%v ", middle)
	}

}

func main() {
	//定义一个数组
	arr := []int{1, 3, 7, 12, 18, 20, 30, 50, 53, 75}
	BinaryFind(&arr, 0, len(arr)-1, 30)
	fmt.Println("main arr=", arr)
}
