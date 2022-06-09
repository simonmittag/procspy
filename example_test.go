package procspy

import (
	"fmt"
)

func Example() {
	lookupProcesses := true
	cs, err := Connections(lookupProcesses)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TCP Connections:\n")
	for c := cs.Next(); c != nil; c = cs.Next() {
		fmt.Printf(" - %v\n", c)
	}
}
