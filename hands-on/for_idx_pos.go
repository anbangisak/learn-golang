// example: https://go.dev/play/p/nGekFavi0S
package main

import "fmt"

func main(){
    // composite literal; slicing literal
    x := []int{7, 9, 42}
    fmt.Println(x)

    y := make([]int, 0, 100) // size setting up
    y = append(y, 7)
    y = append(y, 9)
    y = append(y, 42)
    fmt.Println(y)

    for i, v := range y{
        fmt.Println(i, "-", v)
    }
}
