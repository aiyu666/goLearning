package main

import "fmt"

func main() {
	// var
	var name string = "robin"
	var emptyTypeString = "test"
	var emptyTypeBoolean = true
	var emptyTypeNumber = 0
	fmt.Printf("name: %s", name)
	fmt.Printf("\nemptyTypeString: %s:", emptyTypeString)
	fmt.Printf("\nemptyTypeBoolean: %t", emptyTypeBoolean)
	fmt.Printf("\nemptyTypeNumber: %d", emptyTypeNumber)

	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Printf("\na: %d b: %d c: %d ", a, b, c)

	// :=
	myName, age := "robin", 18
	fmt.Printf("\nmy name is %s and my age is %d", myName, age)

	// const
	const company string = "Trend micro"
	const num_fir, num_sec, num_thir int = 1, 2, 3
	const (
		fir_num   = 3
		sec_num   = 2
		third_num = 1
	)
	fmt.Printf("company: %s", company)
	fmt.Printf("\nnum_fir: %d num_sec: %d num_thir: %d ", num_fir, num_sec, num_thir)
	fmt.Printf("\nfir_num: %d sec_num: %d third_num: %d ", fir_num, sec_num, third_num)
}
