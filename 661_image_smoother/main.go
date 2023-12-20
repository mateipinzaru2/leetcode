package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	result := make([][]int, m)
	for i, row := range img {
		result[i] = make([]int, n)
		for j, val := range row {
			sum := val
			count := 9
			neighbors := [][]int{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
				{i - 1, j - 1},
				{i - 1, j + 1},
				{i + 1, j - 1},
				{i + 1, j + 1},
			}
			for _, neighbor := range neighbors {
				ni, nj := neighbor[0], neighbor[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					sum += img[ni][nj]
				} else {
					count-- // Decrement count for out-of-bounds neighbors
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}

func main() {
	// Prompt for input
	fmt.Println("Enter the dimensions of the image:")
	var m, n int
	fmt.Print("Rows (m): ")
	fmt.Scan(&m)
	fmt.Print("Columns (n): ")
	fmt.Scan(&n)

	// Read CLI input
	img := make([][]int, m)
	fmt.Println("Enter the pixel values:")
	for i := 0; i < m; i++ {
		img[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Printf("Pixel value at (%d, %d): ", i, j)
			fmt.Scan(&img[i][j])
		}
	}

	// Call imageSmoother function
	result := imageSmoother(img)

	// Print the result
	fmt.Printf("Smoothed image: %v\n", result)
}