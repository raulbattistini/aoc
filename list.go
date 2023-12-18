package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
)

type ListNode struct {
	PrevChild *ListNode
	NextChild *ListNode
	Value     interface{}
}
type List struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

var DList List

func (l List) get(i int) interface{} {
	return l.getAt(i).Value
}
func (l *List) getAt(idx int) *ListNode {
	curr := l.Head
	i := 0
	for i < idx {
		curr = curr.NextChild
	}
	return curr
}
func (l *List) append(v interface{}) *List {
	l.Length += 1
	var node ListNode
	node.Value = v
	if l.Tail == nil {
		l.Head = l.Tail
	}
	node.PrevChild = l.Tail
	l.Tail.NextChild = &node
	l.Tail = &node
	return l
}

func hTwo(is []interface{}) *List {
	var l *List
	// get nums, remove 0s
	// l.append([...])
	// extracted value from previous code block when iterating over the string with 0s cleaned up and values like 13 for "1abc3"
	return l
}

func pop(slice []int) ([]int, int, bool) {
	if len(slice) == 0 {
		return slice, 0, false
	}
	element := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, element, true
}
func help(str interface{}) int {
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
			// DList.append(vInt)
			nums = append(nums, vInt)
		}

	}
	if len(nums) == 0 {
		log.Println("length 0")
		// return []int{0}
	} else if len(nums) == 1 {
		log.Println("length 1")
	}
	pattern := `[\[\]0\s]`
	re := regexp.MustCompile(pattern)
	m := re.ReplaceAllString(fmt.Sprint(nums), "")
	if len(m) == 1 {
		m = fmt.Sprintf("%s%s", m, m)
	} else if len(m) > 2 {
		m = fmt.Sprintf("%s%s", m[0:1], m[len(m)-1:])
	}
	val, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println("Error atoi second", err)
		return 0
	}
	fmt.Println("value", val)
	return val
}

func main() {
	strs := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	var f int
	for _, str := range strs {
		v := help(str)
		f += v
	}
	fmt.Println("final", f)
	// func helper(s []string) []int{}

}
