package main

import (
  "fmt"
  "math"
)

func int_len(num int) (int_len int){
  //桁数を知りたい
  if num == 0 {
    int_len = 1
    return int_len
  }
  if num/1 == 0 {
    int_len = 1
    return int_len
  }
  var int_len_pointer int = 1
  int_len = 1
  for {
    int_len_pointer = num / int(math.Pow(10, float64(int_len)))
    if int_len_pointer == 0{
      return int_len
    }
    int_len++
  }
}

func palindromes(num int) (anser_bool bool){
  var i int = 1
  var int_len = int_len(num)
  anser_bool = true
  var num_A int = 0
  var num_B int = 0
  for i <= int_len/2 {
    num_A = num % int(math.Pow(10,float64(i))) / int(math.Pow(10,float64(i)-1.0))
    num_B = num / int(math.Pow(10,float64(int_len-i)))
    if num_A != num_B {
      anser_bool = false
      return anser_bool
    }
    num = num % int(math.Pow(10,float64(int_len-i)))
    i++
  }
  return anser_bool
}

func main(){
  var num int = 0
  var num_A int = 999
  var num_B int = 999
  var anser int = 0
  var anser_A int = 0
  var anser_B int = 0

  for num_A > 0{
    num_B = 999
    for num_B > 0{
      num = num_A * num_B
      if palindromes(num) == true{
        if anser <= num{
          anser = num
          anser_A = num_A
          anser_B = num_B
        }
      }
      num_B = num_B - 1
    }
    num_A = num_A - 1
  }
  fmt.Println(anser,anser_A,anser_B)
}
