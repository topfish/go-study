package main

import (
	"fmt"
)

func main() {
	arr1 := [...]int{1,1,1,2,2,3}
	fmt.Println("arr1:", arr1, "len-arr1:",len(arr1))
	arr2 := make([]int, len(arr1))
	fmt.Println("arr2:", arr2, "len-arr2:", len(arr2))
	var i, j int
		for{
			if arr1[i] == arr1[i+1] {
				if i == j {
					arr2[j] = arr1[1]
				}
				arr2[j] = arr2[j] + arr1[i+1]
				fmt.Println("1-arr2", arr2)
			}else {
				j++
				arr2[j] = arr1[i+1]
				fmt.Println("2-arr2", arr2)
			}
			i++
			fmt.Println("i:", i, len(arr1)-1)
			if i == len(arr1)-1{
				fmt.Println("will be braek....")
				break
			}
		}
	fmt.Println("arr2:", arr2)
	var arr3 []int
	fmt.Println("arr3:", arr3)
	for _, v := range arr2 {
		if v == 0 {
			break
		}
		arr3 = append(arr3, v)
	}
	fmt.Println("---->", arr3)
}
