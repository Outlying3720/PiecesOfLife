package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
    map1 := make(map[int]int)
    map2 := make(map[int]int)
    cnt1 := make([]int, 0)
    cnt2 := make([]int, 0)
    for _,v := range nums1 {
        map1[v] = 1
    }
    for _,v := range nums2 {
        map2[v] = 1
    }

    for k, _ := range map1 {
        if map2[k] != 1 {
            cnt1 = append(cnt1, k)
        }
    }
    for k, _ := range map2 {
        if map1[k] != 1 {
            cnt2 = append(cnt2, k)
        }
    }
    return [][]int{cnt1,cnt2}
}

func main() {
	nums1 := []int{1,2,3,3}
	nums2 := []int{1,1,2,2}
	fmt.Print(findDifference(nums1, nums2))
}



// func getKeys1(m map[int]int) []int {
// 	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
// 	j := 0
// 	keys := make([]int, len(m))
// 	for k := range m {
// 		keys[j] = k
// 		j++
// 	}
// 	return keys
// }

// func findDifference(nums1 []int, nums2 []int) [][]int {
//     map1 := make(map[int]int)
//     map2 := make(map[int]int)
//     cnt1 := make(map[int]int)
//     cnt2 := make(map[int]int)
//     for _,v := range nums1 {
//         map1[v] = 1
//     }
//     for _,v := range nums2 {
//         map2[v] = 1
//     }
//     for _,v := range nums1 {
//         if map2[v] != 1 {
//             cnt1[v] = 1
//         }
//     }
//     for _,v := range nums2 {
//         if map1[v] != 1 {
//             cnt2[v] = 1
//         }
//     }
//     return [][]int{getKeys1(cnt1), getKeys1(cnt2)}
// }