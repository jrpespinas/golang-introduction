package main

import "fmt"

func main() {
	var n, x, y, z int;
	fmt.Scan(&n);
	for i := 1; i <= n; i++ {
		fmt.Scan(&x, &y, &z);
		if x == y && y == z {
			fmt.Printf("Case %v: Equilateral\n", i);
		} else if x == y || y == z || z == x {
			fmt.Printf("Case %v: Isosceles\n", i);
		} else if x + y < z || x + z < y || y + z < x {
			fmt.Printf("Case %v: Invalid\n", i);
		} else {
			fmt.Printf("Case %v: Scalene\n", i);
		}
	}
}
