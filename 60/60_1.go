package main

// import (
// 	"fmt"
// 	"strconv"

// )

// func factorial(n int) int {
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result = result * i
// 	}
// 	return result
// }

// // rel:=""
// // tag:=make(map[int]bool)
// var rel string = ""
// var tag = make(map[int]bool)

// func helper(n, now, k int) string {
// 	sliceNum := factorial(now - 1)
// 	slice := k / sliceNum

// 	if slice*sliceNum == k {
// 		count := 0
// 		t := 1
// 		for ; t <= n; t++ {
// 			if !tag[t] {
// 				count++
// 			}
// 			if count == slice {
// 				break
// 			}
// 		}
// 		rel += strconv.Itoa(t)
// 		tag[t] = true
// 		for i := n; i >= 1; i-- {
// 			if !tag[i] {
// 				rel += strconv.Itoa(i)
// 			}
// 		}
// 	} else {
// 		slice++
// 		count := 0
// 		t := 1
// 		for ; t <= n; t++ {
// 			if !tag[t] {
// 				count++
// 			}
// 			if count == slice {
// 				break
// 			}
// 		}
		
// 		rel += strconv.Itoa(t)
// 		tag[t] = true
// 		helper(n, now-1, k-(slice-1)*sliceNum)
// 	}
// 	return rel
// }

// func getPermutation(n int, k int) string {
// 	if n==0{
// 		return ""
// 	}
// 	helper(n, n, k)
// 	return rel
// }

// func main() {
// 	fmt.Println(getPermutation(9, 1))
// }
