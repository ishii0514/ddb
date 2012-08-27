package db

import (
    "testing"
)


func TestColumnNumber(t *testing.T) {
    var col1 ColumnNumber
    col1.Set("col1")
    if col1.Get() != "col1" {
        t.Error("illegal name.")
    }
}

