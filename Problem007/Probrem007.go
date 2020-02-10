package main

import "fmt"

func prime_num(num int) (prime_num int){
  var i int = 1
  prime := []int{2}

  for {
    FOR_LABEL:
    i = i + 2
    //素数判定
    for _, prime_num := range prime{
      if i % prime_num == 0 {
        //素数じゃない場合
        goto FOR_LABEL
      }
    }
    //nthの素数
    prime = append(prime, i)
    if len(prime) == num {
      prime_num = i
      return prime_num
    }
    goto FOR_LABEL
  }
}

func main(){
  nth := 10001
  tmp := prime_num(nth)
  fmt.Println(tmp)

}
