package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var re1 = regexp.MustCompile(`"[a-z]+"`)
	var str string
	var strs []string
	fmt.Scanln(&str, &str, &str)
	s := re1.FindAllStringSubmatch(str, -1)
	for i := range s {
		strs = append(strs, s[i][0])
	}
	fmt.Println(strs)
	fmt.Println(groupAnagrams(strs))
	// contents = append(contents, s[k][0])

	// numOfString := len(s)
	// if len(s) == 0 {
	// 	fmt.Println(`[[""]]`)
	// 	return
	// }
	// var strs [][]byte
	// for _, v := range s {
	// 	for _, v1 := range v {
	// 		a, _ := strconv.Unquote(v1)
	// 		strs = append(strs, []byte(a))
	// 	}
	// }
	// for i := 0; i < numOfString; i++ {
	// 	var tmp int
	// 	tmp = 1
	// 	for _, v := range strs[i] {

	// 		tmp = tmp * primes[v-97]
	// 	}
	// 	if _, exist := answer[tmp]; !exist {
	// 		answer[tmp] = make([]int, 0)
	// 	}
	// 	answer[tmp] = append(answer[tmp], i)
	// }

	// var output string
	// output = fmt.Sprint(output, "[")
	// for _, v := range answer {
	// 	contents := make([]string, 0)
	// 	for _, k := range v {
	// 		contents = append(contents, s[k][0])
	// 	}
	// 	output = fmt.Sprintf("%s%s,", output, grouping(contents))
	// }
	// output = strings.TrimRight(output, ",")
	// output = fmt.Sprint(output, "]")
	// fmt.Println(output)

}

func groupAnagrams(strs []string) [][]string {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127}
	answer := make(map[int][]int)
	output := make([][]string, 0)

	for i, v := range strs {
		var key int = 1
		v, _ = strconv.Unquote(v)
		for _, t := range []byte(v) {
			key = key * primes[t-97]
		}
		if _, exist := answer[key]; !exist {
			answer[key] = make([]int, 0)
		}
		answer[key] = append(answer[key], i)
	}
	var i int = 0
	for _, v := range answer {
		output = append(output, make([]string, 0))
		for _, item := range v {
			output[i] = append(output[i], strs[item])

		}
		i = i + 1
	}
	return output
}

// func grouping(s []string) string {
// 	var str string
// 	var content string
// 	for _, v := range s {
// 		content = fmt.Sprintf(`%s%s,`, content, v)
// 	}
// 	content = strings.TrimRight(content, ",")
// 	str = fmt.Sprintf("[%s]", content)
// 	return str
// }
