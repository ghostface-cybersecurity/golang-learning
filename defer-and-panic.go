package main
import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "os"
  "log"
)

func enter() float64 {
  reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }

  str = strings.TrimSpace(str)
  number, err := strconv.ParseFloat(str, 64)
  if err != nil {
    log.Fatal(err)
  }
  return number
}
func divide(num_1 float64, num_2 float64) float64 {
  if num_2 == 0 {
    panic("Division by zero !!!!") // allows you to generate an error and exit the program
  } else {
    return (num_1/num_2)
  }
}


func finish(){
  fmt.Println("program has been finished")
}
func ends(){
  fmt.Println("the program ends its execution")
}
func main(){
  defer finish()
  defer ends() // this statement allows the function to be executed at the end of the program where it is called


  fmt.Println("program has been started")


  num_1 := enter()
  num_2 := enter()

  fmt.Println("program is working")

  res := divide(num_1, num_2)
  fmt.Println(res)
}
