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

func reduceHof[I int | string](f func(I, any) any, x any) func(y []I) any {
	return func(z []I) any {
		for _, val := range z {
			x = f(val, x)
		}
		return x
	}
}

func main() {
	my_ints := []int{1, 2, 3, 4, 5}
	my_strings := []string{"Hello", "Big", "Small", "amazing", "traditional", "no"}

	over_three := filterHof(func(i int) bool {
		return i > 3
	})

	add_one := mapHof(func(i int) int {
		return i + 1
	})

	add_world := mapHof(func(s string) string {
		return s + " World!"
	})

	shorter_than_3 := filterHof(func(s string) bool {
		return len(s) < 3
	})

	reduce_count := reduceHof(func(s string, res any) any {
		asserted_num, ok := res.(int)
		if len(s) > 5 {
			if ok {
				asserted_num += 1
			}
		}
		return asserted_num
	}, 0)

	over_three_nums := over_three(my_ints)
	added_nums := add_one(my_ints)
	added_strings := add_world(my_strings)
	shorter_strings := shorter_than_3(my_strings)
	reduced_strings := reduce_count(my_strings)

	log.Println("Numbers over three")
	log.Println(over_three_nums)
	log.Println("Numbers with one added")
	log.Println(added_nums)
	log.Println("Strings with 'world' added")
	log.Println(added_strings)
	log.Println("Strings shorter than three")
	log.Println(shorter_strings)
	log.Println("Count of strings longer than five")
	log.Println(reduced_strings)
}
