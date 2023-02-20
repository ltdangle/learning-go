package main
import "fmt"
type MyInterface interface {
    MyMethod() string
}

type MyStruct struct {
    someField string
}

func (s MyStruct) MyMethod() string {
    return s.someField
}

func main() {
    var myVar MyInterface = MyStruct{"hello world"}

    // Check if myVar implements the MyInterface interface
    if _, ok := myVar.(MyInterface); ok {
        fmt.Println("myVar implements MyInterface")
    } else {
        fmt.Println("myVar does not implement MyInterface")
    }

}

