package main

import "fmt"

func main() {

	fmt.Println("For Loop")
	// First type of for loop
	fmt.Println("First type of for loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Second type of for loop
	fmt.Println("Second type of for loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Third type of for loop - break
	fmt.Println("Third type of for loop - break")
	for {
		fmt.Println("Needs break to come out of loop")
		break
	}

	// Fourth type of for loop - Continue
	fmt.Println("Fourth type of for loop - Continue")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
