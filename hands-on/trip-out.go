package main

import "fmt"

// type is more likely a class in python
type weather string

func (w weather) season() {
    fmt.Println(w)
}

func main(){
    var x weather
    // more like constructor initialization and printing object default values like ORM in dbms django
    x = "rain"
    x.season()
}

