package main

import "fmt"

func main() {
	var x, y int;
	_, err := fmt.Scan(&x, &y);
	for err == nil {
		println(y - x);
		_, err = fmt.Scan(&x, &y);
	}
}
