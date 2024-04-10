package main

type Pos struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	length := m * n
	ans := make([]int, length)
	cur := &Pos{-1, 0}
	turnCnt := 1
	idx := 0
	h := 0
	v := 0
	for idx < length {
		switch turnCnt {
		case 1:
			cur.x += 1
			ans[idx] = matrix[cur.y][cur.x]
			if cur.x == n-1-h {
				turnCnt += 1
			}
		case 2:
			cur.y += 1
			ans[idx] = matrix[cur.y][cur.x]
			if cur.y == m-1-v {
				turnCnt += 1
			}
		case 3:
			cur.x -= 1
			ans[idx] = matrix[cur.y][cur.x]
			if cur.x == h {
				turnCnt += 1
				v += 1
				h += 1
			}
		case 4:
			cur.y -= 1
			ans[idx] = matrix[cur.y][cur.x]
			if cur.y == v {
				turnCnt = 1
			}
		}
		idx += 1
		// fmt.Println(matrix[cur.y][cur.x])
	}
	return ans
}

func main() {
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {9, 10, 11, 12}, {9, 10, 11, 12}}
	spiralOrder(mat)
}
