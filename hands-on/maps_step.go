package main

import "fmt"

func main(){
    x := map[string]int{
        "viva": 52,
        "mac": 61,
    }
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    fmt.Println(x["mac"])

    for k, _ := range x{
        fmt.Println(k, "-", x[k])
    }
}
