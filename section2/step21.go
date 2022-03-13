package main
import (
  "fmt"
  "log"
  "net"
  "io"
)

func main(){
  li, err := net.Listen("tcp", ":8080")
  if(err != nil){
    log.Panic(err)
  }
  defer li.Close()

  for{
    conn, err := li.Accept()
    if(err!=nil){
      log.Panic(err)
    }
    io.WriteString(conn, "\n hello \n")
    fmt.Fprintln(conn, "hai")
    fmt.Fprintln(conn, "hope")
    defer conn.Close()
  }
}
