package _螺旋矩阵__7月22_2刷_

//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//示例 2：
//
//输入：n = 1
//输出：[[1]]
//leetcode:https://leetcode.cn/problems/spiral-matrix-ii/

// 草稿：
// [1,2,3]
// [8,9,4]
// [7,6,5]
//func generateMatrix(n int) [][]int {
//	var a [][]int
//	for i := 0; i < n*n+1; i++ {
//		if i < n {
//			a[0][i] = i + 1
//		}
//		if i == n-1 {
//			a[n][0] = i
//			a[n][1] = i + 1
//		}
//		if 7 > i >= 5 {
//			a[2][1] = i
//			a[2][2] = i + 1
//		}
//		if 9 > i >= 7 {
//			a[0][3] = i
//			a[0][2] = i + 1
//		}
//		//中心点
//		if i%2 == 1 && i == n*n {
//			x := (n - 1) / 2
//			a[x][x] = n * n
//		}
//	}
//	return a
//}

// 题解思路版本
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		//行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列数不变是j 行数变
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变 i 列数变 j--
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变 行变
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
