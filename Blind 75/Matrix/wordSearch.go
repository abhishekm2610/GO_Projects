package main

import (
	"fmt"
)

func exist2(board [][]byte, word string) bool {
	// lookingFor := 0
	rows := len(board)
	cols := len(board[0])

	var dfs func(i int, j int, lookingFor int, visited map[[2]int]int) bool

	dfs = func(i int, j int, lookingFor int, visited map[[2]int]int) bool {

		if i < 0 || j < 0 || i == rows || j == cols || lookingFor == len(word) {

			return false
		}

		if _, p := visited[[2]int{i, j}]; p {

			return false

		}
		if lookingFor == len(word)-1 && board[i][j] == word[lookingFor] {

			return true
		}

		if board[i][j] == word[lookingFor] {
			visited[[2]int{i, j}] = 0
			map1 := make(map[[2]int]int)
			map2 := make(map[[2]int]int)

			map3 := make(map[[2]int]int)

			map4 := make(map[[2]int]int)

			for k, v := range visited {
				map1[k] = v
				map2[k] = v

				map3[k] = v

				map4[k] = v

			}
			r1 := dfs(i+1, j, lookingFor+1, map1)
			r2 := dfs(i-1, j, lookingFor+1, map2)
			r3 := dfs(i, j+1, lookingFor+1, map3)
			r4 := dfs(i, j-1, lookingFor+1, map4)

			return r1 || r2 || r3 || r4

		} else {

			return false
		}

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0, make(map[[2]int]int)) {
				return true
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	// lookingFor := 0
	rows := len(board)
	cols := len(board[0])

	var dfs func(i int, j int, lookingFor int) bool

	dfs = func(i int, j int, lookingFor int) bool {

		if i < 0 || j < 0 || i == rows || j == cols || lookingFor == len(word) {

			return false
		}

		if lookingFor == len(word)-1 && board[i][j] == word[lookingFor] {

			return true
		}

		if board[i][j] == word[lookingFor] {
			fmt.Println(board, string(word[lookingFor]))
			tmp := board[i][j]
			board[i][j] = '1'

			var r1, r2, r3, r4 bool

			r1 = dfs(i+1, j, lookingFor+1)

			r2 = dfs(i-1, j, lookingFor+1)

			r3 = dfs(i, j+1, lookingFor+1)

			r4 = dfs(i, j-1, lookingFor+1)

			board[i][j] = tmp

			return r1 || r2 || r3 || r4

		} else {

			return false
		}

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
func main() {

	nums := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCESEEEFS"
	fmt.Println(exist(nums, word))
}
