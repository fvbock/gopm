package gopm

import "fmt"

const DELIMITER = "--"

type Entry struct {
	Title string
	Text  string
}

func (e *Entry) String() string {
	return fmt.Sprintf("")
}
