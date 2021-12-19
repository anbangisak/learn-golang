package main

import (
    "fmt"
)

func main(){
    x := map[string]int{
        "h": 1,
        "hh": 5,
        "hai": 7,
    }
    fmt.Println(x)

    // can't use for i, key, val := range x { -> i declared but not used error
    // can't use for _, key, val := range x { -> segmentation false -> null pointer dereference
    for key, val := range x {
        fmt.Println(key,"-", val)
    }
}
