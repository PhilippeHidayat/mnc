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
	fmt.Print("Total belanja seorang customer = ")
	input, _ := reader.ReadString('\n')
	inputInt, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Print("Pembeli membayar = ")
	input2, _ := reader.ReadString('\n')
	inputInt2, _ := strconv.Atoi(strings.TrimSpace(input2))

	kembalianAmount, result, success := kembalian(inputInt, inputInt2)
	if success {
		fmt.Printf("Kembalian yang harus diberikan kasir: %d,\n", kembalianAmount)
		kembalianAmount = kembalianAmount - (kembalianAmount % 100)
		fmt.Println("Dibulatkan menjadi: ", kembalianAmount)
		fmt.Println("Pecahan uang:")
		var count = 0
		var curr = 0;
		for _, j := range result {
			if curr != j {
				if(curr != 0) {
					if(curr != 200 && curr != 100) {
						fmt.Printf("%d lembar %d\n", count, curr)
					} else {
						fmt.Printf("%d koin %d\n", count, curr)
					}
					
					count = 0
				}
				curr = j
			} 
			count++
		}
		if(curr != 200 && curr != 100) {
			fmt.Printf("%d lembar %d\n", count, curr)
		} else {
			fmt.Printf("%d koin %d\n", count, curr)
		}
	} else {
		fmt.Println("False, kurang bayar")
	}
}

func kembalian(totalBelanja, bayar int) (int, []int, bool) {
	if bayar < totalBelanja {
		return 0, nil, false
	}
	kembalian := bayar - totalBelanja
	avail := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	var result []int
	for _, v := range avail {
		for kembalian >= v {
			kembalian -= v
			result = append(result, v)
		}
	}
	return bayar - totalBelanja, result, true
}


