package main

import "fmt"
import "math"

const s string = "constant"

func main(){
 fmt.Println("Constant s = ", s)
 const n = 5000
 fmt.Println("sin(5000)",math.Sin(n))

}