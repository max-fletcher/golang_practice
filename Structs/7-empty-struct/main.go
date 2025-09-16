package main

// named empty struct type
type emptyStruct struct{}

func main() {
	// Below are 2 empty structs. They consume 0 bytes i.e no space is consumed at all.

	// anonymous empty struct type
	empty1 := struct{}{}

	empty2 := emptyStruct{}
}
