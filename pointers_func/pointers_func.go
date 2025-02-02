package main
import (
  "fmt"
)

func change_value(x int){
  x = x*x
}

func change_value_pointer(x *int){
  *x = (*x) * (*x)
}

func create_pointer(x int) *int{
  pointer := new(int)
  *pointer = x
  return pointer
}

func main(){
  x := 11
  fmt.Println("---\nchange_value:\n before:",x)
  change_value(x)
  fmt.Println(" after:",x)
  fmt.Println("no changes")
  // ---------
  fmt.Println("---\nchange_value_pointer:\nbefore:",x)
  change_value_pointer(&x)
  fmt.Println("---\n after:",x)
  fmt.Println("success")
  // ----------
  fmt.Println("---\nreturn pointers from funcs")
  p1 := create_pointer(5)
  fmt.Println(p1,*p1)
  p2 := create_pointer(3)
  fmt.Println(p2,*p2)
  p3 := create_pointer(111)
  fmt.Println(p3,*p3)
}
