package main;

func maximalSquare(matrix [][]byte) int {   
	squareSize:= getMaxSquareSizeDP(matrix);
	return squareSize * squareSize;
}

func getMaxSquareSizeDP(matrix [][]byte) int{
	maxSize:= 0;

	var dp = make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	
	for i, row := range matrix{
		for y, cell := range row{
			if y == 0 || i == 0{
				val:= cast(cell);
				dp[i][y] = val; 
				maxSize = max(val, maxSize);
			} else{

				if cell == '1'{
					left:= dp[i][y-1];
					top:= dp[i-1][y];
					topLeft:= dp[i-1][y -1];
	
					neighboursMin := min(min(left, top), topLeft)
					val := neighboursMin + 1;
					dp[i][y] = val;
					maxSize = max(val, maxSize);
				}else{
					dp[i][y] = 0;
				}
			}
		}
	}
	return maxSize;
}

func min(a int, b int) int{
	if a > b {
		return b;
	}
	return a;
}

func cast(cell byte) int{
	if cell == '1'{
		return 1;
	}
	return 0;
}

func getMaxSquareSize(matrix [][]byte) int{
	maxSize:= 0;

	for i, row := range matrix{
		for y, cell := range row{

			if cell == '1'{
				size := squareSearch(matrix, row, i, y, 1);
				maxSize = max(maxSize, size);
			}
		}
	}
	return maxSize;
}

func squareSearch(matrix [][]byte, row []byte, i int, y int, size int) int{
	currentSize := size -1;

	columnCheck := y + currentSize;
	rowCheck := i + currentSize;

	if rowCheck > len(matrix) -1 || columnCheck > len(row) -1{
		return size - 1;
	}

	newColumnMatch := true;
	
	for aux:= 0; aux <= currentSize; aux++{

		val:= matrix[i + aux][columnCheck];
		if val != '1'{
			newColumnMatch = false;
			break;
		}else{
			// matrix[i + aux][columnCheck] = '0';
		}
	}

	rowMatch := true;
	for aux:= 0; aux <= currentSize; aux++{
		val:= matrix[rowCheck][y + aux];
		if val != '1'{
			rowMatch = false;
			break;
		}else{
			// matrix[i + aux][columnCheck] = '0';
		}
	}

	if newColumnMatch && rowMatch{
		return squareSearch(matrix, row, i, y, size + 1);
	}
	return size - 1;
}

func max(a int, b int) int{
	if a > b{
		return a;
	}
	return b;
}
