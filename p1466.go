package main

import "fmt"

// var vis []bool

// func minReorder(n int, connections [][]int) int {
//     e:=make([][]int8, n)
//     for i:=0;i<n;i++ {
//         e[i]=make([]int8,n+1)
//     }
//     vis=make([]bool, n)

//     for i:=0;i<len(connections);i++ {
//         a:=connections[i][0]
//         b:=connections[i][1]
//         e[a][b] = -1
//         e[b][a] = 1
//     }
//     vis[0] = true
//     return dfs(0, e)
// }

// func dfs(cur int, e [][]int8) int {
//     cost:=0
//     for i:=0; i<len(e); i++ {
//         if vis[i] == false {
//             if e[cur][i] == -1 {
//                 vis[i] = true
//                 cost = cost + 1 + dfs(i, e)
//                 fmt.Println(i,"->",cur," +1")
//             } else if e[cur][i] == 1 {
//                 vis[i] = true
//                 cost = cost + dfs(i, e)
//                 fmt.Println(i,"->",cur," ")
//             }
//         }
//     }
//     return cost
// }

func minReorder(n int, connections [][]int) int {
    e:=make([][][]int, n)
    for _,item := range connections {
        a:=item[0]
        b:=item[1]
        e[a] = append(e[a], []int{b, 1})
        e[b] = append(e[b], []int{a, 0})
    }
    
    var res int

    var dfs func(int, int)
    dfs = func(cur, root int) {
        for _, next := range e[cur] {
            if next[0] == root {
                fmt.Println(cur,"->",next[0]," root")
                continue
            }
            if next[1] == 1{
                res += 1
                fmt.Println(cur,"->",next[0]," +1")
            }
            dfs(next[0], cur)
        }
    }
    dfs(0, -1)
    return res
}

func main() {
	conn:=[][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}
	n := 6
	fmt.Println(minReorder(n,conn))
}