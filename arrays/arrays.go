package main
import (
  "fmt"
)

func main(){
  // array init

  var arr[5]int
  for i:=0; i<5; i++{
    fmt.Println(i)
    arr[i] = i
  }
  fmt.Println(arr)

  fmt.Println("---")
  // or

  var arr2[5]int = [5]int{1,2,3,4,5}
  fmt.Println(arr2)

  fmt.Println("---")
  // or

  nums := [5]int{1,2,3,4,5}
  fmt.Println(nums)

  fmt.Println("---")

  // if an ellipsis is specified instead of an array, then the size is determined based on the number of elements
  nums2 := [...]int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(nums2)
  for i:=0; i<10; i++{
    fmt.Println("[",i+1,"] -",nums2[i])
  }

  fmt.Println("---")
}
