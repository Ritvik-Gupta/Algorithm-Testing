package easy

func generate(numRows int) [][]int {

	triangle := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		triangle = append(triangle, make([]int, i+1))
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}
