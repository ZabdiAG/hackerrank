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
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {
	chars := []byte(w)
	changed := false
	for i := len(chars) - 1; i >= 1; i-- {
		if chars[i] > chars[i-1] {
			changed = true
			// find next highest than chars[i-1]
			highest := chars[i]
			idxHighest := i
			for x := i; x < len(chars); x++ {
				if chars[x] < highest && chars[x] > chars[i-1] {
					highest = chars[x]
					idxHighest = x
				}
			}
			chars[i-1], chars[idxHighest] = chars[idxHighest], chars[i-1]
			// order  chars[i+1:] from lowest to highest
			for j := i; j < len(chars); j++ {
				for k := j; k < len(chars); k++ {
					if chars[j] > chars[k] {
						chars[j], chars[k] = chars[k], chars[j]
					}
				}
			}

			break
		}
	}
	if changed {
		return string(chars)
	}
	return "no answer"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
