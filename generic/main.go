package generic

type Sortable interface {
	Len() int
	Compare(idx1 int, idx2 int) int
	Swap(idx1 int, idx2 int)
}

func Sort(arr Sortable) {
	for i := 0; i < arr.Len()-1; i++ {
		for j := i + 1; j < arr.Len(); j++ {
			if arr.Compare(i, j) > 0 {
				arr.Swap(i, j)
			}
		}
	}
}

type IntArray []int

func (this IntArray) Len() int {
	return len(this)
}

func (this IntArray) Compare(idx1, idx2 int) int {
	return this[idx1] - this[idx2]
}

func (this IntArray) Swap(idx1, idx2 int) {
	tmp := this[idx1]
	this[idx1] = this[idx2]
	this[idx2] = tmp
}

type StringArray []string

func (this StringArray) Len() int {
	return len(this)
}

func (this StringArray) Compare(idx1, idx2 int) int {
	return len(this[idx1]) - len(this[idx2])
}

func (this StringArray) Swap(idx1, idx2 int) {
	tmp := this[idx1]
	this[idx1] = this[idx2]
	this[idx2] = tmp
}
