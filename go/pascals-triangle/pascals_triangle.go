// Package pascal returns the Pascal's Triangle
package pascal

// Triangle returns the Pascal's Triangle for the given level that is asked for, layers numbered from 1
func Triangle(level int) [][]int {
	level--

	// first layer is free
	triangle := [][]int{{1}}

	for n := 1; n <= level; n++ {
		// start counting from 1 because first row is free and build each row

		triangle = append(triangle, []int{1}) // start with a 1 on the left side of every layer
		for i := 0; i < len(triangle[n-1])-1; i++ {
			triangle[n] = append(triangle[n], triangle[n-1][i]+triangle[n-1][i+1])
		}
		triangle[n] = append(triangle[n], 1) // finish with a 1 on the right side of every layer
	}

	return triangle
}