package mockx_test

import (
	"errors"
	"fmt"

	"github.com/AndreyArthur/mockx"
)

// Define a sorting interface.
type Sorting interface {
	IsSorted(slice []int) (bool, error)
}

// Implement a Searcher struct that depends on the Sorting interface.
type Searcher struct {
	sorting Sorting
}

func NewSearcher(sorting Sorting) *Searcher {
	return &Searcher{
		sorting: sorting,
	}
}

func (searcher *Searcher) linear(slice []int, target int) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1
}
func (searcher *Searcher) binary(slice []int, target int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func (searcher *Searcher) Search(slice []int, target int) int {
	sorted, err := searcher.sorting.IsSorted(slice)
	if err != nil {
		panic(err)
	}

	if sorted {
		return searcher.binary(slice, target)
	}
	return searcher.linear(slice, target)
}

// Define a Sorting mock struct.
type SortingMock struct {
	mockx.Mockx
}

func NewSortingMock() *SortingMock {
	sorting := &SortingMock{}
	sorting.Init((*Sorting)(nil))

	return sorting
}

func (sorting *SortingMock) IsSorted(slice []int) (bool, error) {
	values := sorting.Call("IsSorted", slice)
	return mockx.Value[bool](values[0]), mockx.Reference[error](values[1])
}

// Use the mockx library in tests.
func Example_usage() {
	sorting := NewSortingMock()
	searcher := NewSearcher(sorting)

	slice := []int{3, 1, 2, 5, 4}

	sorting.Return("IsSorted", false, nil)
	index := searcher.Search(slice, 3)
	fmt.Printf("index = %d\n", index)

	sorting.Return("IsSorted", true, nil)
	index = searcher.Search(slice, 3)
	fmt.Printf("index = %d\n", index)

	sorting.Impl("IsSorted", func(slice []int) (bool, error) {
		if slice == nil {
			return false, errors.New("Cannot verify, a nil slice was given.")
		}

		for i := range len(slice) - 1 {
			if slice[i] > slice[i+1] {
				return false, nil
			}
		}
		return true, nil
	})
	index = searcher.Search(slice, 4)
	fmt.Printf("index = %d\n", index)

	// Output:
	// index = 0
	// index = -1
	// index = 4
}
