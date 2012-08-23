package main 

import(
    "fmt"
    "db"
)

func main() {
    
    var col1 db.ColumnNumber
    col1.Set("col1")
    fmt.Println(col1.Get()) 
    
    fmt.Println(db.Get())
    
}

