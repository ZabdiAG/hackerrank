package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Node represents graph node
type Node struct {
	N     uint32
	Edges []*Node
}

type Graph struct {
	Start   Node
	Current Node
}

func (*Graph) GetNode(N uint32) *Node {

}

/*
 * Complete the jennysSubtrees function below.
 */
func jennysSubtrees(n int32, r int32, edges [][]int32) int32 {
	/*
	 * Write your code here.
	 */
	fmt.Printf("r: %v\n", r)
	mainNode := Node{1}
	for n := range egges {

	}
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nr := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var edges [][]int32
	for edgesRowItr := 0; edgesRowItr < int(n-1); edgesRowItr++ {
		edgesRowTemp := strings.Split(readLine(reader), " ")

		var edgesRow []int32
		for _, edgesRowItem := range edgesRowTemp {
			edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
			checkError(err)
			edgesItem := int32(edgesItemTemp)
			edgesRow = append(edgesRow, edgesItem)
		}

		if len(edgesRow) != int(2) {
			panic("Bad input")
		}

		edges = append(edges, edgesRow)
	}

	result := jennysSubtrees(n, r, edges)

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
