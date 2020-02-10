package main

import "fmt"

func int_array(start int, end int) []int{
  var int_array []int
  for i := start; i <= end; i++{
    int_array = append(int_array, i)
  }
  return int_array
}

func mult_sum(num_array []int) (mult_sum_num int){
  for _, i := range num_array {
    mult_sum_num = mult_sum_num + i*i
  }
  return mult_sum_num
}

func sum_mult(num_array []int) (sum_mult_num int){
  for _, i := range num_array{
    sum_mult_num = sum_mult_num + i
  }
  sum_mult_num = sum_mult_num * sum_mult_num
  return sum_mult_num
}


func main(){
  var num_array []int
  num_array = int_array(1, 100)
  anser := sum_mult(num_array) - mult_sum(num_array)
  fmt.Println(anser)
}
