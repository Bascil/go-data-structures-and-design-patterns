package main

import "fmt"

func main(){
	nums := []int{1,2,3,4,5,5,6}
	//fmt.Println(nums[0])

	// count the number of elements
	//println(len(nums))

	// iterate over slice
	for index, value := range nums {
		println("Index", index,"Value",value)
	}

	nums = append(nums, 100,120,150)
	fmt.Println(nums)


	// maps
	user := map[string]string{
		"name": "Basil",
		"age": "35",
		"role": "developer",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	names, ok := user["names"]
	if ok {
		fmt.Println(names)
	} else {
		println("Invalid key")
	}

	user["names"] = "people"
	fmt.Println(user)

	for k, v := range user {
		println(k, v)
	}
}