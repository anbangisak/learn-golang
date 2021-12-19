package main

import "fmt"

type colors interface{
    rgb() string
}

type bw struct{
    black string
    white string
}

func (b bw) rgb() string {
    return b.black
}

type rg struct{
    red string
    green string
}

func (r rg) rgb() string {
    return r.red
}

// simple package funtion
func sayColor(c colors){
    fmt.Println("color is ", c.rgb())
}

func main(){
    x := bw{"forest", "snow"}
    y := rg{"rose", "leaves"}

    sayColor(x)
    sayColor(y)
}
