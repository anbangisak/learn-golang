// methods receivers make methods

package main

import "fmt"

type color struct {
    white string
    black string
}

type person struct{
    fname string
    lname string
}

// func (receiver) identifier(parameters) (returns) { <code> }
// receiver -> self for class in python or this method in c++ class
func (c color) speak() {
    fmt.Println(c.white, "near", c.black)
}

func (p person) speak() {
    fmt.Println(p.fname, "-", p.lname)
}

func main(){
    x := color{
        "snow",
        "forest",
    }
    x.speak()

    y := person{
        "thikey",
        "uchiha",
    }
    y.speak()
}
