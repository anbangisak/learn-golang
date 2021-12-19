package main

import "fmt"

func main(){
    fmt.Println(checking())
    x := checking()
    fmt.Println(x)
}

// func (receiver) identifier(parameter) (returns) { <code> }
func checking() int{
    return 5
}
