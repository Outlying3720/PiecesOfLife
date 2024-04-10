package gameoflife

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Game struct {
	m     int
	n     int
	board [][]int
}

func NewGame(board [][]int) *Game {
	// 00 死 01 活变死 10 死变活 11 活
	g := &Game{}
	g.m = len(board)
	g.n = len(board[0])
	g.board = board
	return g
}

func (g Game) PrintFrame() {
	optSys := runtime.GOOS
	if optSys == "linux" {
		//执行clear指令清除控制台
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		//执行clear指令清除控制台
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	for y := 0; y < g.m; y++ {
		for x := 0; x < g.n; x++ {
			fmt.Print(g.board[y][x], " ")
		}
		fmt.Print("\n")
	}
}

func (g Game) Next() {
	for y := 0; y < g.m; y++ {
		for x := 0; x < g.n; x++ {
			g.board[y][x] += g.nextStatus(x, y)
		}
	}
	for y := 0; y < g.m; y++ {
		for x := 0; x < g.n; x++ {
			g.board[y][x] = g.board[y][x] >> 1
		}
	}
}

func (g Game) nextStatus(x, y int) int {
	alive := 0
	dirs := [][]int{{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1}}
	for _, dir := range dirs {
		cx := x + dir[0]
		cy := y + dir[1]
		if cx < 0 || cy < 0 || cx > g.n-1 || cy > g.m-1 {
			continue
		}
		if g.board[cy][cx]&1 == 1 {
			alive += 1
		}
	}
	if alive < 2 {
		return 0
	}
	if alive == 2 && g.board[y][x]&1 == 1 {
		return 2
	}
	if alive == 3 {
		return 2
	}
	return 0 // >3
}
