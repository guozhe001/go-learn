package main

import (
	"fmt"
)

func main() {
	initAllSquares()
	fmt.Println(allSquares)
	// result := numSquares(12)
	// fmt.Println("result=", result)
}

// 所有的完全平方数记录下来
var allSquares []int

// int最大值
var max int = int(^uint(0) >> 1)

// 使用map保存完全平方数，key为完全平方数，value为0
var quareMap map[int]int = make(map[int]int)

func initAllSquares() {
	index := 0
	number := 1
	for number < 100000 {
		quare := number * number
		if quare > max {
			break
		}
		quareMap[quare] = 0
		allSquares = append(allSquares, quare)
		fmt.Println(quare)
		index++
		number++
	}
}

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {

	// 最少的是多少呢？
	// 如果这个数字本身是一个可以开方的数字，则结果是1，是最小
	// 否则，获取比这个值小的那个最大的完全平方数
	return 0
}

// @lc code=end
