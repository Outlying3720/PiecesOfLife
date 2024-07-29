package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	i := 0
	n := len(gas)
	for i < n {
		j := i
		tank := gas[j]
		cnt := 0
		fmt.Println(j, tank, cnt, i)
		for cnt < n && tank >= cost[j] {
			tank -= cost[j]
			j = (j + 1) % n
			tank += gas[j]
			cnt++
			fmt.Println(j, cnt, tank)
		}
		if cnt == n {
			return i
		}
		if i > j {
			break
		}
		i = j + 1
	}

	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	// a:=0
	// fmt.Println(a++)
	fmt.Println(canCompleteCircuit(gas, cost))
}
