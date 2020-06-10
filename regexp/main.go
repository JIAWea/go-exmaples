package main

import (
    "fmt"
    "regexp"
)

func main()  {
	s1 := "/of_course_i_still_love_you"
	s2 := "/of_course_i_still_love_you/000"
	reg := regexp.MustCompile(`/of_course_i_still_love_you/`)
	
	fmt.Println(reg.FindAllString(s1, -1)) // []
	fmt.Println(reg.FindAllString(s2, -1)) // [/of_course_i_still_love_you/]

	ok, _ := regexp.MatchString(`/of_course_i_still_love_you/\d+`, s1)
	fmt.Println("match: ", ok) // false
	ok, _ = regexp.MatchString(`/of_course_i_still_love_you/\d+`, s2)
	fmt.Println("match: ", ok) // true
}
