package main

import "fmt"

//二分查找

func main(){
	matrix := [][]int {{1,3,5,7},{10,11,16,20},{23,30,34,50}}

	target1 := 3
	ret1 := searchMatrix(matrix,target1)
	fmt.Println("ret1 = ", ret1)

	target2 := 13
	ret2 := searchMatrix(matrix,target2)
	fmt.Println("ret2 = ", ret2)
}


func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}
	if len(matrix[0]) <= 0 {
		return false
	}

	seq_start := Seq2DimTo1Dim(0, 0, len(matrix[0]))
	seq_end := Seq2DimTo1Dim(len(matrix)-1, len(matrix[0])-1, len(matrix[0]))

	for(seq_start <= seq_end) {
		seq_mid := (seq_start+seq_end)/2
		row_mid, column_mid := Seq1DimTo2Dim(seq_mid, len(matrix[0]))

		if matrix[row_mid][column_mid] == target {
			return true
		} else if matrix[row_mid][column_mid] > target {
			seq_end = seq_mid-1
		} else {
			seq_start = seq_mid+1
		}
	}
	return false
}

func Seq2DimTo1Dim(row, column, rowLen int) int {
	return row * rowLen + column
}

func Seq1DimTo2Dim(seq, rowLen int) (int, int) {
	return seq/rowLen, seq%rowLen
}