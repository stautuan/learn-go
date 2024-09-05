/*
We support various visualization dashboards on Textio that display message 
analytics to our users. The UI for our graphs and charts is built on top of
a grid system. Let's build some grid logic.

Complete a function that takes a number of rows and colums and returns a 2D
slice of integers. The value of each cell is `i * j` where `i` and `j` are
the indexes of the row and column respectively.
*/

package main

// create a 2D slice of integers with the given number of rows and columns
// each cell is i * j, where i and j are indexes of row and column
func createMatrix(rows, cols int) [][]int {
    // initialize a 2D slice
    matrix := make([][]int, rows)
    // iterate over each row
    for i := 0; i < rows; i++ {
        // each row creates a new slice of integers with length cols
        matrix[i] = make([]int, cols)
        // iterate over each column in the current row
        for j := 0; j < cols; j++ {
            matrix[i][j] = i * j
        }
    }
    return matrix
   
}
