package sort

type Sortee interface {
	GetLength() int
	Get(index int) int
	Swap(index1 int, index2 int)
	Return() interface{}
}

type arraySortee struct {
	array []int
}

func newArraySortee(array []int) *arraySortee {
	return &arraySortee{array: array}
}

func (sortee *arraySortee) GetLength() int {
	return len(sortee.array)
}

func (sortee *arraySortee) Get(index int) int {
	return sortee.array[index]
}

func (sortee *arraySortee) Swap(index1 int, index2 int) {
	sortee.array[index1], sortee.array[index2] = sortee.array[index2], sortee.array[index1]
}

func (sortee *arraySortee) Return() interface{} {
	return sortee.array
}

type twoDArraySortee struct {
	twoDimArray [][]int
	colIndex int
}

func new2DArraySortee(array [][]int) *twoDArraySortee {
	return &twoDArraySortee{
		twoDimArray: array,
		colIndex:    -1,
	}
}

func (sortee *twoDArraySortee) MoveToCol(index int) {
	sortee.colIndex = index
}

func (sortee *twoDArraySortee) GetLength() int {
	return len(sortee.twoDimArray)
}

func (sortee *twoDArraySortee) Get(index int) int {
	return sortee.twoDimArray[index][sortee.colIndex]
}

func (sortee *twoDArraySortee) Swap(index1 int, index2 int) {
	sortee.twoDimArray[index1], sortee.twoDimArray[index2] = sortee.twoDimArray[index2], sortee.twoDimArray[index1]
}

func (sortee *twoDArraySortee) Return() interface{} {
	return sortee.twoDimArray
}


