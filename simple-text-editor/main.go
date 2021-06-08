package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// OperationIdentifier identifier to recognize change type on text
type OperationIdentifier int

// OperationIdentifiers
const (
	OperationAppend OperationIdentifier = iota + 1
	OperationDelete
	OperationPrint
	OperationUndo
)

// Change ...
type Change struct {
	Type OperationIdentifier
	Diff []byte
}

// TextEditor ...
type TextEditor struct {
	Text    []byte
	History []Change
}

// Append writes w to the end of text
func (te *TextEditor) Append(w []byte) {
	te.History = append(te.History, Change{OperationAppend, w})
	te.Text = append(te.Text, w...)
}

// Delete removes k characters at the end of the text
func (te *TextEditor) Delete(k int) {
	diff := make([]byte, k)
	copy(diff, te.Text[len(te.Text)-k:])
	te.History = append(te.History, Change{OperationDelete, diff})
	te.Text = te.Text[:len(te.Text)-k]
}

// Print character at k index
func (te *TextEditor) Print(k int) {
	fmt.Println(string(te.Text[k-1]))
}

// Undo last operation(Append || Delete)
func (te *TextEditor) Undo() {
	lastChange := te.History[len(te.History)-1]
	switch lastChange.Type {
	case OperationAppend:
		te.Text = te.Text[:len(te.Text)-len(lastChange.Diff)]
	case OperationDelete:
		te.Text = append(te.Text, lastChange.Diff...)
	}
	te.History = te.History[:len(te.History)-1]
}

func main() {
	var textEditor = TextEditor{
		History: make([]Change, 0),
	}
	var err error
	var Q int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if Q, err = strconv.ParseInt(scanner.Text(), 10, 64); err != nil {
			panic(err)
		}
	}
	for i := int64(1); scanner.Scan() && i <= Q; i++ {
		q := scanner.Text()
		ops := strings.Split(q, " ")
		op, err := strconv.ParseInt(ops[0], 10, 32)
		if err != nil {
			panic(err)
		}
		switch OperationIdentifier(op) {
		case OperationAppend:
			textEditor.Append([]byte(ops[1]))
		case OperationDelete:
			v, err := strconv.ParseInt(ops[1], 10, 32)
			if err != nil {
				panic(err)
			}
			textEditor.Delete(int(v))
		case OperationPrint:
			v, err := strconv.ParseInt(ops[1], 10, 32)
			if err != nil {
				panic(err)
			}
			textEditor.Print(int(v))
		case OperationUndo:
			textEditor.Undo()
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
