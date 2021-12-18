package main

import "fmt"

type person struct{
    fname string
    lname string
}

func main(){
    v := 7
    fmt.Println(v)
    fmt.Printf("%T\n", v)

    bb := []string{"hai", "hello"}
    fmt.Println(bb[0])

    ba := map[string]int{"ooi": 8, "fresh": 9, "meat": 0}
    fmt.Println(ba["fresh"])

    bc := person{
        "gisak",
        "uchiha",
    }
    fmt.Println(bc.lname, bc.fname)
}
