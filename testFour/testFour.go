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
	fmt.Print("Jumlah cuti bersama = ")
	input, _ := reader.ReadString('\n')
	inputInt, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Print("Tanggal bergabung = ")
	input2, _ := reader.ReadString('\n')
	fmt.Print("Tanggal cuti = ")
	input3, _ := reader.ReadString('\n')
	fmt.Print("Durasi cuti = ")
	input4, _ := reader.ReadString('\n')
	inputInt4, _ := strconv.Atoi(strings.TrimSpace(input4))
	canTake, message := canTakePrivateLeave(inputInt, input2, input3, inputInt4)
	fmt.Println(message)
	if canTake {
		fmt.Println("True")
	}
}

func canTakePrivateLeave(commonLeave int, joinDate string, plannedLeave string, leaveDuration int) (bool, string) {
	privateLeave := 14 - commonLeave

	year, _, _ := getYearMonthDay(plannedLeave)
	endYear := year
	endMonth := 12
	endDay := 31
	totalDays := countDays(joinDate, endYear, endMonth, endDay)
	privateLeaveThisYear := int((float64(totalDays - 180)/365)*float64(privateLeave))
	if totalDays < 180 {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	// check if private leave quota is enough
	if privateLeaveThisYear < leaveDuration {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", privateLeaveThisYear)
	}

	// check if private leave is max 3 days in a row
	if leaveDuration > 3 {
		return false, "Cuti pribadi max 3 hari berturutan"
	}

	return true, ""
}

func countDays(startDate string, endYear int, endMonth int, endDay int) int {
	startYear, startMonth, startDay := getYearMonthDay(startDate)
	days := 0
	for year := startYear; year <= endYear; year++ {
		for month := 1; month <= 12; month++ {
			if year == startYear && month < startMonth {
				continue
			}
			if year == endYear && month > endMonth {
				break
			}
			for day := 1; day <= 31; day++ {
				if year == startYear && month == startMonth && day < startDay {
					continue
				}
				if year == endYear && month == endMonth && day > endDay {
					break
				}
				days++
			}
		}
	}
	return days
}

func getYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	return year
}

func getMonth(date string) int {
	month, _ := strconv.Atoi(date[5:7])
	return month
}

func getDay(date string) int {
	day, _ := strconv.Atoi(date[8:10])
	return day
}

func getYearMonthDay(date string) (int, int, int) {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	return year, month, day
}
