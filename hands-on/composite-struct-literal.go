package main

import "fmt"

// struct declartion doesn't need comma for line separation
// only values need comma separation
type color struct{
    red int
    blue int
}

func main(){
    // package struct usage
    x := color{255, 8}
    fmt.Println(x)

    // block struct creation on execution
    y := struct{
        white string
        black string
    }{
        "snow",
        "forest",
    }

    fmt.Println(y)
}
