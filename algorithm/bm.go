package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

func BMCheck(context string, comparison string)  (int, int){
	c := Search(context, comparison)
	if c == -1 {
		c = c + 1
		fmt.Println("失敗")
	}else if c == 0{
		fmt.Println("失敗")
	}else {
		fmt.Println("成功")
		value := context[:c]
		minutevalue := OutNumValue(value)
		return c, minutevalue
	}
	return c, 0
}

func Search(haystack, needle string) (int) {
	c := strings.Index(haystack, needle)
	fmt.Println(c)
	return c
}

func OutNumValue(str string) int {
	minutevalue, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Atoiに失敗しとりますぜ")
		fmt.Println(err)
		return 0
	}else {
		fmt.Println(minutevalue)
	}
	return minutevalue
}

