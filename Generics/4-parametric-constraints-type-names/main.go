package main

// You can name the type anything instead of using "T" as shown below. Though using T is the common convention, it is good to know since you might
// be confused with T's everywhere.
func splitAnySlice1[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitAnySlice2[MyAnyType any](s []MyAnyType) ([]MyAnyType, []MyAnyType) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

type customer interface {
	GetBillingEmail() string
}

// example using an interface that isn't "any"
func splitAnySlice3[MyAnyType2 customer](s MyAnyType2) string {
	return s.GetBillingEmail()
}
