package main

import "fmt"

func Value(n int, m int, arr []int) int {
	//điểm dừng cho đệ quy
	if n == 1 {
		if arr[0] <= m {
			return 1
		}
		return 0
	}
	// chia nhỏ giá trị với số ẩn giảm dần
	count := 0
	for i := 1; i <= m; i++ {
		if arr[n-1]*i <= m {
			count += Value(n-1, m-arr[n-1]*i, arr)
		}
	}
	return count
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var arr []int
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	value := Value(n, m, arr)
	fmt.Println(value)

}
