// map = has table
package main
import "fmt"

func main(){
  var words map[string]int // init
  characters = map[string]int{
    "s" : 5,
    "b" : 6,
    "c" : 7,
  }
  fmt.Println(people)
  // add elements
  var words = map[string]int{ // keys - string ; values - int
    "Milk" : 1,
    "Steak" : 2,
    "Pencil" : 3,
    "BRO" : 5,
    ":)" : 8,
    // thank you my brothers :)
  }
  fmt.Println(peoples[":)"])
  fmt.Println(peoples["Milk"])
  peoples["Milk"] = 111
  fmt.Println(peoples["Milk"])
  // checking values
  if val, ok := peoples["BRO"]; ok{
    fmt.Println(val, ok)
  }
  // enumeration of elements
  for key, value := range peoples{
    fmt.Println(key, value)
  }
  // The make function provides an alternative option for creating a mapping. It creates an empty hash table:
  make_results := make(map[int]int)
  // add and delete elements
  // add
  make_results[5] = 3
  make_results[7] = 10
  for i := 10; i < 20 ; i++{
    make_results[i] = i+20
  }
  fmt.Println(make_results)
  // delete
  delete(make_results, 5)
  for i := 10; i < 20; i++{
    delete(make_results, i)
  } 
  fmt.Println(make_results)
}
