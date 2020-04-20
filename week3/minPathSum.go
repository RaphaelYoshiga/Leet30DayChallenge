package main;

type Point struct {
	x, y int
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0{
		return 0;
	}

	cells := make(map[Point]int);

	checkPath(0, 0, grid, 0, cells);

	return cells[Point{x: len(grid) -1, y:len(grid[0]) -1}]
}

func checkPath(x int, y int, grid [][]int, sum int, cells map[Point]int){
	
	rowLen := len(grid);
	colLen := len(grid[0]);

	if x < 0 || y < 0 || x >= rowLen || y >= colLen{
		return;
	}

	point:= Point{x: x, y: y};
	sum+= grid[x][y];

	if cellSum, found := cells[point]; found{
		if cellSum <= sum{
		   return;
		}else{
			cells[point] = sum;
		}
   } else {
		cells[point] = sum;
   }

	// checkPath(x - 1, y, grid, sum, cells); //Up
	checkPath(x + 1, y, grid, sum, cells); //Down
	// checkPath(x, y -1, grid, sum, cells); // Left
	checkPath(x, y + 1, grid, sum, cells); //Right
}