package main

import "fmt"
import "strconv"

func fizzBuzz(num int) {
  if num == 0 {
    return
  }
  if num % 3 ==0 && num % 5 == 0 {
      fmt.Print("FizzBuzz\n")
  } else if num % 5 == 0 {
      fmt.Print("Buzz\n")
  } else if num % 3 == 0 {
    fmt.Print("Fizz\n")
  } else {
    numStr :=  strconv.Itoa(num)
    fmt.Print(numStr + "\n")
  }
  fizzBuzz(num - 1)
}

func main() {
  fizzBuzz(100)
}
