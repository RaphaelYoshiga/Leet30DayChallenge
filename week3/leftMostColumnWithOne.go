package main;
type BinaryMatrix struct {
	// Get(x int, y int) int
	Dimensions() []int
}
/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

 func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions();

	rowLen := dimensions[0];
	colLen := dimensions[1];

	valSet:=false;
	minIndex := colLen -1;
	for r:= 0; r < rowLen; r++{
		for c:= minIndex; c >= 0; c--{
			
			val:= binaryMatrix.Get(r,c);

			if val == 1{
				minIndex = min(minIndex, c)
				valSet:= true;
			}else{
				break;
			}
		}
	}

	if valSet {
		return minIndex;
	}

    return -1;
}

func min(a int, b int) int{
	if a > b {
		return b;
	}
	return a;
}