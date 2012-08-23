package db

import (
    "testing"
)

func TestXYZ(t *testing.T) {
    var col1 db.ColumnNumber
    col1.Set("col1")
    fmt.Println(col1.Get()) 
    Equal(t, "col1",col1.Get(), "illegal name!")
    t.Error("show")
}

