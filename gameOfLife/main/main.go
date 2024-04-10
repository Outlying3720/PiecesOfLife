package main

import (
	"crypto/rand"
	gameoflife "gameoflife/GameOfLife"
	"time"
)

func main() {
	board := [][]int{}
	for i := 0; i < 100; i++ {
		row := [200]int{}
		for j := 0; j < 100; j++ {
			row[j] = rand.Int()
		}
	}

	game := gameoflife.NewGame(board)
	for {
		game.PrintFrame()
		game.Next()
		time.Sleep(time.Second)
	}
}
