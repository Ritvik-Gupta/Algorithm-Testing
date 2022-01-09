package main

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		rowElms, columnElms := make(map[int]struct{}), make(map[int]struct{})
		for j := 0; j < n; j++ {
			if _, ok := rowElms[matrix[i][j]]; ok {
				return false
			}
			if _, ok := columnElms[matrix[j][i]]; ok {
				return false
			}
			rowElms[matrix[i][j]] = struct{}{}
			columnElms[matrix[j][i]] = struct{}{}
		}
	}
	return true
}
