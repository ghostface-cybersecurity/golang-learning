package main
import (
  "fmt"
)

func main(){
  // -----------
  fmt.Println("---")
  var num int = 5
  var p *int
  p = &num
  fmt.Println("Number:",num)
  fmt.Println("Pointers:\n p:",p,"\n *p", *p)
  // -----------
  fmt.Println("---")
  fmt.Println("*p = 25")
  *p = 25
  fmt.Println("Number:",num,"\nPointers:\n p:",p,"\n *p",*p)
  // ------------
  fmt.Println("---")
  f := 47
  pf := &f
  fmt.Println("pf - pointer f\n f:",f,"\n pf:",pf,"\n *pf",*pf)
  // ------------
  fmt.Println("---")
  *pf = 52
  fmt.Println("*pf = 52\n *pf:",*pf,"\n pf:",pf)
  // ------------
  fmt.Println("---\nEmpty pointer")
  var nil_pointer *float64
  if nil_pointer != nil {
    fmt.Println("nil_pointer:",*nil_pointer)
  }
  // -----------
  fmt.Println("---\nthe new function")
  pointer := new(int)
  fmt.Println("New pointer:",*pointer,"|",pointer) // default - 0
  *pointer = 77
  fmt.Println("New pointer:",*pointer,"|",pointer) // now - 77
}
