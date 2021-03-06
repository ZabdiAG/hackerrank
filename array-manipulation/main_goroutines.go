package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

// TODO:
// 10/16 test cases failed, pasing: [0, 1, 2, 3, 14, 15]
// 4. OUTPUT: 7540746475, ( variant)
// 7. TimeLimit exceded, however, output is correct: 2497169732
func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	arr := make([]int64, n)
	arrPointers := make([]*int64, n)
	for i := range arr {
		arrPointers[i] = &arr[i]
	}
	maxNumber := int64(0)
	chEnd := make(chan bool)
	for _, v := range queries {
		// Loop on arr starting from a to b, inclusive
		go func(chEnd chan bool, arr []*int64, maxNumber *int64, v []int32) {
			for it := v[0] - 1; it <= (v[1] - 1); it++ {
				*arr[it] = *arr[it] + int64(v[2])
				if *arr[it] > *maxNumber {
					*maxNumber = *arr[it]
				}
			}
			chEnd <- true
		}(chEnd, arrPointers, &maxNumber, v)
	}
	counter := 0
	for range chEnd {
		if counter < len(queries)-1 {
			counter++
			continue
		}
		close(chEnd)
		counter++
	}
	return maxNumber
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
