package main
import (
  "fmt"
)

func main(){
  // slice init
  var users []string
  fmt.Println(users,"\n---OR---")
  // or
  var users1 = []string{"Bob", "Alice"}
  fmt.Println(users1,"\n---OR---")
  // or
  users2 := []string{"1","2","3","NoName"}
  fmt.Println(users2,"\n---OR---")
  // ---
  // Slice elements can be accessed by index
  for _, value := range users2{
    fmt.Println(value)
  }
  for i := 2; i < 2; i++{
    fmt.Println(users[i])
  }
  fmt.Println("---")
  // ---
  // another way to create a slice with default values
  fmt.Println("Another way to create a slice with default values")
  var values []int = make([]int, 5)
  for i := 0; i<2; i++{
    values[i] = i
  }
  // Add to slice
  fmt.Println("---\nAdd to slice")
  values = append(values, 5)
  for i := 9 ; i < 11; i++{
    values = append(values,i)
  }
  fmt.Println(values)
  // slice operators
  fmt.Println("---\nslice operators")
  init_vals := []int{1,2,3,4,5,6,7,8}
  vals1 := init_vals[0:2] // 0 - 3
  vals2 := init_vals[2:4] // 4 - 5
  vals3 := init_vals[4:] // 5 - end
  fmt.Println(vals1)
  fmt.Println(vals2)
  fmt.Println(vals3)
  // ---
  // delete elements from slice
  // delete 5 element
  fmt.Println("---\ndelete elements from slice")
  fmt.Println(init_vals)
  n := 4
  new_vals := append(init_vals[:n],init_vals[n+1:]...) // three dots - used to unpack the slice when passing it to the append function
  fmt.Println(new_vals)
}
