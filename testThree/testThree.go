package main

import "fmt"

func main() {
	var str = "[<{<{[{[{}[[<[<{{[<[<[[[<{{[<<<[[[<[<{{[<<{{<{<{<[<{[{{[{{{{[<<{{{<{[{[[[{<<<[{[<{<<>>[]}]>>>}]]]}]}>}}}>>]}}}}]}}]}>]>}>}>}}>>]}}>]>]]]>>>]}}>]]]>]>]}}>]>]]]}]}>}>]"
	fmt.Println(isValid(str))
}

func isValid(str string) bool {
	stack := make([]rune, 0)
	allowed := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}
	for _, c := range str {
		if _, ok := allowed[c]; ok {
			stack = append(stack, c)
		} else if len(stack) == 0 || allowed[stack[len(stack)-1]] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
