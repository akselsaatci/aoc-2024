package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("First answer : %d\nSecond answer : %d", firstAnswer(), secondAnswer())

}

func secondAnswer() int {

	clueFile, err := os.Open("./clue.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	map1 := make(map[int]int)
	map2 := make(map[int]int)

	scanner := bufio.NewScanner(clueFile)
	scanner.Split(bufio.ScanWords)
	isItFirst := true
	for scanner.Scan() {
		elem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf(err.Error())
		}
		if isItFirst {
			map1[elem] += 1
		} else {
			map2[elem] += 1
		}
		isItFirst = !isItFirst
	}
	res := 0

	for key := range map1 {
		times, ok := map2[key]
		if !ok {
			continue
		}

		res += times * key
	}

	return res

}

func firstAnswer() int {
	clueFile, err := os.Open("./clue.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	var arr1, arr2 []int
	scanner := bufio.NewScanner(clueFile)
	scanner.Split(bufio.ScanWords)
	isItFirst := true
	for scanner.Scan() {
		elem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf(err.Error())
		}
		if isItFirst {
			arr1 = append(arr1, elem)
		} else {
			arr2 = append(arr2, elem)
		}
		isItFirst = !isItFirst
	}
	res := 0
	sort.Ints(arr1)
	sort.Ints(arr2)

	for index, elem := range arr1 {
		currentRes := elem - arr2[index]

		if currentRes < 0 {
			currentRes = -currentRes
		}

		res += currentRes
	}

	return res

}
