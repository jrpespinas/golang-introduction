package main

import "fmt"

func main() {
	var n, x int;
	var med, max int;
	fmt.Scan(&n);
	for i := 1; i <= n; i++ {
		max = 0;
		med = 0;
		for j := 0; j < 3; j++ {
			fmt.Scan(&x);
			if x > max {
				med = max;
				max = x;
			} else if x > med {
				med = x;
			}
		}
		fmt.Printf("Case %v: %v\n", i, med);
	}
}
