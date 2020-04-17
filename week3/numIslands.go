package main;

func numIslands(grid [][]byte) int {

	c := 0;

	for x, row := range grid{
		for y, cell := range row{
			if cell == '1'{
				c++;
				lookAllDirections(x,y, grid);
			}
			
		}
	}

	return c;
}

func lookAllDirections(x int,y int, grid [][]byte){
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return;
	}

	grid[x][y] = '0';

	lookAllDirections(x-1, y, grid);
	lookAllDirections(x+1, y, grid);
	lookAllDirections(x, y-1, grid);
	lookAllDirections(x, y+1, grid);
}

func main(){

	data := [][]byte{ 
		{'1','1','1'},
		{'0','1','0'},
		{'1','1','1'},
	};

	r:= numIslands(data);
	r++;
}