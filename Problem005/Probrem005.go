package main

import "fmt"

func int_array(start int, end int) []int{
  var int_array []int
  for i := start; i <= end; i++{
    int_array = append(int_array, i)
  }
  return int_array
}

func smallest_multiple(num int, num_array []int) bool {
  var anser_bool bool = true
  for _, i := range num_array{
    if num % i != 0{
      anser_bool = false
    }
  }
  return anser_bool
}

func main(){
  var num_array []int
  var i int = 0
  num_array = int_array(1, 20)
  for {
    i++
    if smallest_multiple(i, num_array) == true{
      break
    }
  }
  fmt.Println(i)
}
