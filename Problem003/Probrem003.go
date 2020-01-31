package main

import "fmt"

func factorization(jude_number int) []int{
  var anser []int
  var i int = 1
  prime := []int{2}

  if jude_number % 2 == 0 {
    anser = append(anser, 2)
  }
  for jude_number != 1 {
    FOR_LABEL:
    i = i + 2
    //素数判定
    for _, prime_num := range prime{
      if i % prime_num == 0 {
        goto FOR_LABEL
      }
    }
    //素因数の判定式
    prime = append(prime, i)
    if jude_number % i == 0 {
      anser = append(anser, i)
      jude_number = jude_number / i
      fmt.Println("%d",jude_number)
    }
  }
  return anser
}

func main(){
  var JUDE_NUMBER int = 600851475143
  tmp := factorization(JUDE_NUMBER)
  fmt.Println(tmp)

}
