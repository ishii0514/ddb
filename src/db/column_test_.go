package db

import (
    "testing"
)


func TestColumnInteger(t *testing.T) {
    var col1 ColumnInteger
    col1.SetName("col1")
    if col1.GetName() != "col1" {
        t.Error("illegal name.")
    }
}

func TestColumnIntegerInsert(t *testing.T) {
    var col1 ColumnInteger
    col1.Insert(1)
    if col1.DataCount() != 1 {
        t.Error("illegal datacount.")
    }
    col1.Insert(1)
    if col1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
}
func TestColumnIntegerGet(t *testing.T) {
    var col1 ColumnInteger
    col1.Insert(2)
    col1.Insert(1)
    if col1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
    
    val,err := col1.Get(0)
    if err != nil {
        t.Error("illegal data get error.")
    }
    if val != 2 {
        t.Error("illegal data get.")
    }
    
    //範囲外の指定
    val,err = col1.Get(10)
    if err == nil {
        t.Error("illegal error.")
    }
    if err.Error() != "out of range." {
        t.Error("illegal error message.")
    }
}
func TestColumnIntegerDeleteAt(t *testing.T) {
    var col1 ColumnInteger
    col1.Insert(2)
    col1.Insert(1)
    col1.DeleteAt(1)
    if col1.DataCount() != 1 {
        t.Error("illegal delete.")
    }
    
    val,err := col1.Get(0)
    if err != nil {
        t.Error("illegal data get error.")
    }
    if val != 2 {
        t.Error("illegal data delete.")
    }
}