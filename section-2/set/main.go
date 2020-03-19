package main

import "fmt"

type Set map[string]struct{}

func main() {
	s := make(Set)

	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k := range s {
		retVal = append(retVal, k)

	}
	return retVal
}
