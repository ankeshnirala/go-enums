package main

import (
	"fmt"
	"reflect"
)

type Status int

const (
	Delete Status = iota + 1
	Active
	Blocked
)

func (s Status) String() string {
	return [...]string{"Delete", "Active", "Blocked"}[s-1]
}

func main() {
	var status = Delete
	fmt.Println(reflect.TypeOf(status.String()))
}
