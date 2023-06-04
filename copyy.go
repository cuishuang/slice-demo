package main

func main() {
	var s = []int{1, 2, 3}

	for i, n := range s {
		if i == 0 {
			s[1], s[2] = 8, 9
		}
		print(n)
	}
}
