package main

// Apart from interfaces, we can use this syntax to set multiple types.

// Here, Ordered matches any type that supports <, <=, >, and >=. SO it won't throw a compile-time error.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Because T is constrained by Ordered, the compiler knows
// that < is valid for any T used with this function.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
