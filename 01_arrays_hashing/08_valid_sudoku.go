package arrays_hashing_01

/*
Valid Sudoku
Solved
You are given a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:

Each row must contain the digits 1-9 without duplicates.
Each column must contain the digits 1-9 without duplicates.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
Return true if the Sudoku board is valid, otherwise return false

Note: A board does not need to be full or be solvable to be valid.

Example 1:

Input: board =
[["1","2",".",".","3",".",".",".","."],

	["4",".",".","5",".",".",".",".","."],
	[".","9","8",".",".",".",".",".","3"],
	["5",".",".",".","6",".",".",".","4"],
	[".",".",".","8",".","3",".",".","5"],
	["7",".",".",".","2",".",".",".","6"],
	[".",".",".",".",".",".","2",".","."],
	[".",".",".","4","1","9",".",".","8"],
	[".",".",".",".","8",".",".","7","9"]]

Output: true
Example 2:

Input: board =
[["1","2",".",".","3",".",".",".","."],

	["4",".",".","5",".",".",".",".","."],
	[".","9","1",".",".",".",".",".","3"],
	["5",".",".",".","6",".",".",".","4"],
	[".",".",".","8",".","3",".",".","5"],
	["7",".",".",".","2",".",".",".","6"],
	[".",".",".",".",".",".","2",".","."],
	[".",".",".","4","1","9",".",".","8"],
	[".",".",".",".","8",".",".","7","9"]]

Output: false
Explanation: There are two 1's in the top-left 3x3 sub-box.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/
func isValidSudoku(board [][]byte) bool {

	check := func(row []byte) bool {
		set := make(map[byte]struct{}, len(row))
		for _, b := range row {
			if string(b) == "." {
				continue
			}
			_, exists := set[b]
			if exists {
				return false
			}
			set[b] = struct{}{}
		}
		return true
	}

	for i := range board {
		if !check(board[i]) {
			return false
		}
	}

	for i := range board[0] {
		column := make([]byte, len(board), len(board))
		for j := 0; j < len(board); j++ {
			column[j] = board[j][i]
		}
		if !check(column) {
			return false
		}
	}

	for i := 0; i < 3; i++ { // horizontal
		for j := 0; j < 3; j++ { // vertical

			q := []byte{
				board[0+i*3][0+j*3], board[0+i*3][1+j*3], board[0+i*3][2+j*3],
				board[1+i*3][0+j*3], board[1+i*3][1+j*3], board[1+i*3][2+j*3],
				board[2+i*3][0+j*3], board[2+i*3][1+j*3], board[2+i*3][2+j*3],
			}
			if !check(q) {
				return false
			}

		}
	}

	return true
}

/*
Recommended Time & Space Complexity
You should aim for a solution as good or better than O(n^2) time and O(n^2) space, where n is the number of rows in the
square grid.


Hint 1
Which data structure would you prefer to use for checking duplicates?


Hint 2
You can use a hash set for every row and column to check duplicates. But how can you efficiently check for the squares?


Hint 3
We can find the index of each square by the equation (row / 3) * 3 + (col / 3). Then we use hash set for O(1) lookups
while inserting the number into its row, column and square it belongs to. We use separate hash maps for rows, columns,
and squares.
*/
