package main

import "fmt"

func main(){
  var fib_a int = 1
  var fib_b int = 2
  var fib_tmp int = 0
  var anser int = 2
  for anser <= 4000000 {
    fib_tmp = fib_a + fib_b
    if fib_tmp % 2 == 0 {
      anser = anser + fib_tmp
    }
    fib_a = fib_b
    fib_b = fib_tmp
  }
  fmt.Println(anser)
}
