package db

import (
    "testing"
)


func TestColumnInteger(t *testing.T) {
    var col1 = ColumnInteger{name : "col1"}
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
func TestColumnIntegerSearch(t *testing.T) {
    var col1 ColumnInteger
    col1.Insert(1)
    col1.Insert(2)
    col1.Insert(1)
    res := col1.Search(1)
    if len(res) != 2 {
        t.Error("illegal result len.")
    }
    if res[0] != 0 {
        t.Error("illegal search.#1")
    }
    if res[1] != 2 {
        t.Error("illegal search. #2")
    }
}
func TestColumnIntegerSearchNoMatch(t *testing.T) {
    var col1 ColumnInteger
    col1.Insert(1)
    col1.Insert(2)
    col1.Insert(1)
    res := col1.Search(3)
    if len(res) != 0 {
        t.Error("illegal result len.")
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