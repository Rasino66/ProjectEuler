package main

import "fmt"

func main(){
  var anser int = 0
  for i := 0; i<1000; i++{
    if i % 3 == 0 {
      anser = anser +i
    } else if i % 5 == 0 {
      anser = anser + i
    }
  }
  fmt.Println(anser)
}
