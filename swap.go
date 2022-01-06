package main

import "fmt"

func main() {
	a := 10
  b := 20
  var t int = a
  a = b
  b = t
  fmt.Println("Value of a: ",a)
  fmt.Println("Value of b: ",b)

}
