package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func flatten(nested []interface{}) []interface{} {
	flattened := make([]interface{}, 0)

	for _, i := range nested {
		switch i.(type) {
		case []interface{}:
			flattenedSubArray := flatten(i.([]interface{}))
			flattened = append(flattened, flattenedSubArray...)
		case interface{}:
			flattened = append(flattened, i)
		}
	}

	return flattened
}

func helper(str interface{}) []int {
	var n uint8 = 1
	var s string = ""
	var nums []int
	for i := range str.(string) {
		if reflect.TypeOf(str.(string)[i]) == reflect.TypeOf(s) {
			continue
		}
		if reflect.TypeOf(str.(string)[i]) == reflect.TypeOf(n) {
			val := str.(string)[i]
			b := []byte{val}
			vInt, _ := strconv.Atoi(string(b))
			// if err != nil {
			//         fmt.Println("Error converting number", err)
			// }
			nums = append(nums, vInt)
		}

	}
	if len(nums) == 0 {
		log.Println("length 0")
		return []int{0}
	} else if len(nums) == 1 {
		nums = append(nums, nums[0])
	}
	fmt.Println("nums", nums)
	return nums
}

// func dayone(strs []string) int {
//         var allNumsFromStrs []interface{}
//         var f []interface{}
//         var t int
//         for _, str := range strs {
//                 numsFromStr := helper(str)
//                 allNumsFromStrs = append(allNumsFromStrs, numsFromStr)
//         }
//         for i := range allNumsFromStrs {
//                 fint := allNumsFromStrs[i].([]int)
//                 expr := fmt.Sprintf("%d,", fint[i])
//                 f = append(f, expr)
//         }
//         // f = [[arr1] [arr2]]
//
//         fmt.Println("f t", f, t)
//         return t
// }

func dayone(strs []string) int {
	nStrs := []int{}
	var t int
	var r int
	for _, str := range strs {
		rect := helper(str)
		for j := range rect {
			r += rect[j]
		}
		fmt.Println("r", r)
		nStrs = append(nStrs, r)
	}
	for i := range nStrs {
		t += nStrs[i]
	}
	return t
}

func mainer() {
	strs := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	right := dayone(strs)
	fmt.Println("vim-go", right)
}
