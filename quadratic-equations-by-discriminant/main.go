package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
  "strings"
  "strconv"
  "math"
)

func GetString(namevar string) string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter ",namevar,":")
  str, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  return str
}


func FloatConvert(str string, namevar string) float64 {
  fmt.Println("[*] Convertation of variable |",namevar,"|")
  str = strings.TrimSpace(str)
  number, err := strconv.ParseFloat(str,64)
  if err != nil {
    log.Fatal(err)
  }
  return number
}


func main(){
  fmt.Println("[Hello] Welcome to discriminant programm :)")

  a := GetString("a")

  b := GetString("b")

  c := GetString("c")

  fmt.Println("[INFO]\n  a =", a, "\n  b =", b,"\n  c =",c)
  fmt.Println("[+] Success\n---")

  // --------------

  fmt.Println("[INFO] Convert strings to number...")

  a_float := FloatConvert(a, "a")

  b_float := FloatConvert(b, "b")

  c_float := FloatConvert(c, "c")

  fmt.Println("[+] Success\n---")

  D := (b_float * b_float) - (4 * a_float * c_float)

  x1 := (-b_float + math.Sqrt(D)) / (2.0 * a_float)
  x2 := (-b_float - math.Sqrt(D)) / (2.0 * a_float)
  fmt.Println(" x1:",x1)
  fmt.Println(" x2:",x2)
}
