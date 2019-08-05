package main

func getAverage(arr [5] int, size int) float64 {
	var sum float64
	for i := 0; i < size; i++ {
		sum += float64(arr[i])
	}
	return sum / float64(size)
}

func main() {
	var balance = [5] int {1000, 2, 3, 17, 50}
	println(getAverage(balance, 5))
}
