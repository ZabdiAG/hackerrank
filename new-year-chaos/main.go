package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// TODO: end testcases not passing due to timeout

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	bribes := 0
	qbribes := make([]int, len(q))
	for i := 1; i < len(q); i++ {
		for j := 0; j < len(q)-i; j++ {
			if q[j] > q[j+1] {
				qbribes[q[j]-1]++
				if qbribes[q[j]-1] > 2 || qbribes[q[j+1]-1] > 2 {
					fmt.Printf("Too chaotic\n")
					return
				}
				q[j], q[j+1] = q[j+1], q[j]
				bribes++
			}
		}
	}
	fmt.Printf("%d\n", bribes)
}

/*
func minimumBribes(q []int32) {
    // 1 2 3 4 5 6 7 8
    // 1 2 5 3 7 8 6 4
    bribes := 0
    for i := 0 ; i < len(q) ; i++ {
        if q[i]-1 == int32(i) {
            continue
        }
        diff := q[i] - (int32(i) +1)
        if diff < 0 {
            continue
        }
        if diff > 2 {
            fmt.Printf("Too chaotic\n")
            return
        }
        bribes+=int(diff)
    }
    fmt.Printf("%d\n", bribes)
}
*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
