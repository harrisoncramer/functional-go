package main

import "log"

type Ints []int
type Lists [][]int
type Strings []string

// Filter expects a function that will be evaluated against each item, and filter out those that return false.
func filterHof[I int | string](f func(x I) bool) func(y []I) []I {
	return func(z []I) []I {
		res := []I(nil)
		for _, val := range z {
			if f(val) {
				res = append(res, val)
			}
		}
		return res
	}
}

// Map expects a function that will be run against each item, and return the same type, but modified.
func mapHof[I int | string](f func(x I) I) func(y []I) []I {
	return func(z []I) []I {
		res := []I(nil)
		for _, val := range z {
			res = append(res, f(val))
		}
		return res
	}
}

func main() {
	my_ints := []int{1, 2, 3, 4, 5}
	my_strings := []string{"Hello", "Big", "Small"}

	over_five := filterHof(func(i int) bool {
		return i > 3
	})

	add_one := mapHof(func(i int) int {
		return i + 1
	})

	add_world := mapHof(func(s string) string {
		return s + " World!"
	})

	over_five_nums := over_five(my_ints)
	added_nums := add_one(my_ints)
	added_strings := add_world(my_strings)

	log.Println(over_five_nums)
	log.Println(added_nums)
	log.Println(added_strings)
}
