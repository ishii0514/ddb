package db

import (
    "testing"
)


func TestColumnNumber(t *testing.T) {
    var col1 ColumnNumber
    col1.SetName("col1")
    if col1.GetName() != "col1" {
        t.Error("illegal name.")
    }
}

func TestColumnNumberInsert(t *testing.T) {
    var col1 ColumnNumber
    col1.Insert(1)
    if col1.DataCount() != 1 {
        t.Error("illegal datacount.")
    }
    col1.Insert(1)
    if col1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
}
func TestColumnNumberGet(t *testing.T) {
    var col1 ColumnNumber
    col1.Insert(2)
    col1.Insert(1)
    var val = col1.Get(0)
    if val != 2 {
        t.Error("illegal data get.")
    }
    if col1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
}
func TestColumnNumberDelete(t *testing.T) {
    var col1 ColumnNumber
    col1.Insert(2)
    
    col1.Insert(1)
//    col1.Delete(2)
    if col1.DataCount() != 1 {
        t.Error("illegal datacount.")
    }
}