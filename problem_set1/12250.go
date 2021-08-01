package main

import "fmt"

func main() {
	var i int = 1;
	var lang string;
	for true {
		fmt.Scanln(&lang);
		if lang == "#" {
			break;
		}

		if lang == "HELLO" {
			fmt.Printf("Case %v: %v\n", i, "ENGLISH");
		} else if lang == "HOLA" {
			fmt.Printf("Case %v: %v\n", i, "SPANISH");
		} else if lang == "HALLO" {
			fmt.Printf("Case %v: %v\n", i, "GERMAN");
		} else if lang == "BONJOUR" {
			fmt.Printf("Case %v: %v\n", i, "FRENCH");
		} else if lang == "CIAO" {
			fmt.Printf("Case %v: %v\n", i, "ITALIAN");
		} else if lang == "ZDRAVSTVUJTE" {
			fmt.Printf("Case %v: %v\n", i, "RUSSIAN");
		} else {
			fmt.Printf("Case %v: %v\n", i, "UNKNOWN");
		}
		i++;
	}
}
