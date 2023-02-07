package main

import "fmt"

func main()  {
    a1:=[3]int{1,2,3} 
    a2:=[3]int{4,5,6} 
    slice:=append(a1[:len(a1)],a2[:len(a2)]...)
    fmt.Println(slice)
}
