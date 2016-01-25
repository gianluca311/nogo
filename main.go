package main
import (
    "fmt"
    "github.com/gianluca311/nogo/notam"
)

func main() {
    fmt.Println("fetching now for LOWW")
    notamList := notam.GetNotams("LOWW")
    for _,elem := range notamList.Notamlist {
        fmt.Println(elem)
    }
}