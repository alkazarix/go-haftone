package halftone

type Filter struct {
	Name   string
	Matrix [][]int32
}

func newFilter(name string, matrix [][]int32) *Filter {
	return &Filter{
		Name:   name,
		Matrix: matrix,
	}
}

func (f Filter) Min() int32 {
	m := f.Matrix
	min := m[0][0]
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] < min {
				min = m[x][y]
			}
		}
	}
	return min
}

func (f Filter) Max() int32 {
	m := f.Matrix
	max := m[0][0]
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] > max {
				max = m[x][y]
			}
		}
	}
	return max
}

var Bayer = newFilter(
	"bayer",
	[][]int32{
		[]int32{0, 2},
		[]int32{3, 1},
	},
)

var Bayer2 = newFilter(
	"bayer2",
	[][]int32{
		[]int32{0, 8, 2, 10},
		[]int32{12, 4, 14, 6},
		[]int32{3, 11, 1, 9},
		[]int32{15, 7, 13, 5},
	},
)

var Bayer4 = newFilter(
	"bayer4",
	[][]int32{
		[]int32{0, 48, 12, 60, 3, 61, 15, 63},
		[]int32{32, 16, 44, 28, 35, 19, 47, 21},
		[]int32{8, 56, 4, 52, 11, 59, 7, 55},
		{40, 24, 36, 20, 43, 27, 39, 23},
		{2, 50, 14, 62, 1, 49, 13, 61},
		{34, 18, 46, 30, 33, 17, 45, 29},
		{10, 58, 6, 54, 9, 57, 5, 53},
		{42, 26, 38, 22, 41, 25, 37, 21},
	},
)
