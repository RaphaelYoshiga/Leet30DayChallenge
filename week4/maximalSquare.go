package main;

func maximalSquare(matrix [][]byte) int {   
	squareSize:= getMaxSquareSize(matrix);
	return squareSize * squareSize;
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

func main(){
	bytes:= [][]byte {
        {'0','1','1','0','0','1','0','1','0','1'},
        {'0','0','1','0','1','0','1','0','1','0'},
        {'1','0','0','0','0','1','0','1','1','0'},
        {'0','1','1','1','1','1','1','0','1','0'},
        {'0','0','1','1','1','1','1','1','1','0'},
        {'1','1','0','1','0','1','1','1','1','0'},
        {'0','0','0','1','1','0','0','0','1','0'},
        {'1','1','0','1','1','0','0','1','1','1'},
        {'0','1','0','1','1','0','1','0','1','1'}};
	maximalSquare(bytes);
}