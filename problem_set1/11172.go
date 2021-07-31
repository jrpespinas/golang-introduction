package main

import "fmt"

func main() {
	var n, x, y int;
	fmt.Scan(&n);
	for n > 0 {
		fmt.Scan(&x, &y);
		if x > y {
			println(">");
		} else if x < y {
			println("<");
		} else {
			println("=");
		}
		n--;
	}
}
