package main

import "fmt"

func main() {
	a_main()
}

/*func test() {
	var array [101][101]int
	array = fillingArray(array, 1, 1)
	cnt := countingArray(array)
	fmt.Println(cnt)
}*/

func a_main() {
	var array [101][101]int
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ { //n의 갯수만큼 좌표값 입력받기
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		array = fillingArray(array, a, b) //입력받은 좌표에서부터 10*10에 1을 채움
	}
	fmt.Println(countingArray(array))
}
func fillingArray(array [101][101]int, a, b int) [101][101]int { //원래 배열과 a,b 좌표를 받아 좌표에서부터 10*10만큼 배열에 1을 채움
	for i := a; i < a+10; i++ {
		for j := b; j < b+10; j++ {
			array[i][j] = 1
		}
	}
	return array
}

func countingArray(array [101][101]int) int { //원래배열을 받아 배열 내의 1 갯수를 리턴
	var cnt int = 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if array[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}
