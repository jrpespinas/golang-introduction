package main

import "fmt"

func main() {
	var n, x, y, z int;
	var med int;
	fmt.Scan(&n);
	for i := 1; i <= n; i++ {
		fmt.Scan(&x, &y, &z);
		if x > y && x > z {
			if y > z {
				med = y;
			} else {
				med = z;
			}
		} else if y > z && y > x {
			if x > z {
				med = x;
			} else {
				med = z;
			}
		} else if x > y {
			med = x;
		} else {
			med = y;
		}
		fmt.Printf("Case %v: %v\n", i, med);
	}
}
