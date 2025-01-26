package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputInt, _ := strconv.Atoi(strings.TrimSpace(input))
	var array []string
	array = make([]string, inputInt)
	for i := 0; i < inputInt; i++ {
		content, _ := reader.ReadString('\n')
		array[i] = strings.TrimSpace(content)
	}
	duplicate(array)
}

func duplicate(array []string) {
	result := make(map[string][]int)
	var duplicateKey = ""
	for i := 0; i < len(array); i++ {
		array[i] = strings.ToLower(array[i])
		if _, ok := result[array[i]]; ok {
			if duplicateKey == "" {
				duplicateKey = array[i]
			}
			
			result[array[i]] = append(result[array[i]], i + 1)
		} else {
			result[array[i]] = []int{i + 1}
		}
	}
	fmt.Println(duplicateKey)
	if duplicateKey != "" {
	value := result[duplicateKey]
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(value), " ", " ", -1), "[]"))
	} else {
		fmt.Println("false")
	}
}