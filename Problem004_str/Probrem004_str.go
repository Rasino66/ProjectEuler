package main

import (
  "fmt"
  "github.com/grohiro/xrange"
)

func main() {
  for i := xrange.NewIntRange(100, 999){
    for j := xrange.NewIntRange(100,999){
      fmt.Println(i,j)
    }
  }
}
