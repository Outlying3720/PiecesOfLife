package main

import "fmt"

type pos struct {
    x, y int
}

var dir = []pos{{-1,0}, {1,0}, {0,-1}, {0,1}}

var m int
var n int
var lstr int

func dfs(matrix [][]byte, str string, cnt int, curpos pos) bool {
	if cnt == lstr {
		return true
	}
	tmp := matrix[curpos.x][curpos.y]
	matrix[curpos.x][curpos.y] = '+'
	for di:=0; di<len(dir); di++ {
		d := dir[di]
		if curpos.x + d.x < m && curpos.x + d.x >= 0 && curpos.y + d.y < n && curpos.y + d.y >= 0 {
			if matrix[curpos.x + d.x][curpos.y + d.y] == str[cnt] {
				// tmp := matrix[curpos.x + d.x][curpos.y + d.y]
				// matrix[curpos.x + d.x][curpos.y + d.y] = '+'
				// fmt.Print(matrix[curpos.x + d.x][curpos.y + d.y], " ")
				done := dfs(matrix, str, cnt+1, pos{curpos.x + d.x, curpos.y + d.y})
				// matrix[curpos.x + d.x][curpos.y + d.y] = tmp
				if done {
					return true
				}
			}
		}
	}
	matrix[curpos.x][curpos.y] = tmp
	return false
}

func hasPath(matrix [][]byte, str string) bool {
    m=len(matrix)
	if m == 0 {
		if len(str) == 0 {
			return true
		}
		return false
	}
    n=len(matrix[0])
	lstr = len(str)
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if matrix[i][j] == str[0] {
				curpos := pos{i, j}
                if dfs(matrix, str, 1, curpos) {
					return true
				}
            }
        }
    }

    return false
}

func main() {
	matrix := [][]byte{{'A','A'}}
	str := "AAA"

	fmt.Println(hasPath(matrix, str))
}