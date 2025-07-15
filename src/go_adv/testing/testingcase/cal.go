package testingcase

// 用于测试的函数
func add(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}
