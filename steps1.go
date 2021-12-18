package main

import "fmt"

type human interface{
    speak()
}

func saySomething(h human){
    h.speak()
}

type person struct{
    fname string
    lname string
}

func (p person) speak(){
    fmt.Println(p.lname, p.fname)
}

type secretAgent struct{
    person
    allowedToKill bool
}

func (sa secretAgent) speak(){
    fmt.Println(sa.fname,sa.lname,sa.allowedToKill)
}

func main(){
    v := 7
    fmt.Println(v)
    fmt.Printf("%T\n", v)

    bb := person{
        "uchiha",
        "thikey",
    }

    fmt.Println(bb)

    sa1 := secretAgent{
        bb,
        true,
    }
    fmt.Println(sa1)

    saySomething(bb)
    saySomething(sa1)

}
