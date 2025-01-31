package main
import (
  "fmt"
  "os"
  "bufio"
  "log"
  "strings"
  "strconv"
)

func enter() int {
  reader := bufio.NewReader(os.Stdin)
  number_str, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }

  number_str = strings.TrimSpace(number_str)
  number, err := strconv.Atoi(number_str)
  if err != nil {
    log.Fatal(err)
  }
  return number
}

func factorial(a int) int {
  if a == 0{
    return 1
  } else {
    return a * factorial(a-1)
  }
}

func main(){
  fmt.Println("Enter a number:")
  num := enter()
  res := factorial(num)
  fmt.Println(res)
}
