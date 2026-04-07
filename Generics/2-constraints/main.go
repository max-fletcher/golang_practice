package main

// The commented interface below is used when we refer the "any" type for generics. It is an empty interface that doesn't have any constraints.
// The problem is that using "any" might cause unexpected runtime errors. (See example below).
// type any interface {
// }

// This interface makes sure that the struct passed has a String function associated to it. The name "toString" can be changed to anything.
// This will prevent any struct that doesn't have a function called "String" to be rejected(type error), hence preventing runtime errors.
type toString interface {
	String() string
}

func concat[T toString](vals []T) string {
	result := ""
	for _, val := range vals {
		// this is where the .String() method
		// is used. That's why we need a more specific
		// constraint instead of the any constraint
		result += val.String()
	}
	return result
}
