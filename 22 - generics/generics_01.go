package main 

import "fmt" 

func PrintAnything[T any](s []T) { 
  for _, v := range s { 
    fmt.Println(v) 
  } 
} 

func PrintInt(v int) { 
  println(v) 
} 

func main() { 
  PrintAnything([]string{"Hello, ", "World"}) 
  PrintAnything([]int{4, 2}) 
  PrintAnything([]float64{3.7, 2.9}) 
  PrintInt(7) 
}