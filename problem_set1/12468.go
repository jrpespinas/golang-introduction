package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return x * -1;
	} else {
		return x;
	}
}

func min(x, y int) int {
	if x < y {
		return x;
	} else {
		return y;
	}
}
func main() {
	var curr, next, min1, min2 int;
	for true {
		fmt.Scan(&curr, &next);
		if curr == -1 && next == -1 {
			break;
		}

		min1 = abs((100 - curr) - next);
		min2 = abs(next - curr);

		println(min(min1, min2));
	}
}
