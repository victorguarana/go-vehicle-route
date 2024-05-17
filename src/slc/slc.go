// This package contains functions to manipulate slices.
package slc

func CircularSelection[T any](list []T, index int) T {
	i := index % len(list)
	return list[i]
}

func Copy[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func InsertAt[T any](slice []T, element T, index int) []T {
	newSlice := make([]T, len(slice)+1)
	for i := 0; i < len(newSlice); i++ {
		switch {
		case i < index:
			newSlice[i] = slice[i]
		case i == index:
			newSlice[i] = element
		case i > index:
			newSlice[i] = slice[i-1]
		}
	}
	return newSlice
}

func RemoveElement[T comparable](points []T, point T) []T {
	newPoints := make([]T, 0)
	for _, p := range points {
		if p != point {
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}