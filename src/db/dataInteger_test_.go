package db

import (
    "testing"
)

func TestArrayIntegerInsert(t *testing.T) {
    var data1  = ArrayInteger{}
    data1.Insert(1)
    if data1.DataCount() != 1 {
        t.Error("illegal datacount.")
    }
    data1.Insert(1)
    if data1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
}

func TestArrayIntegerGet(t *testing.T) {
    var data1  = ArrayInteger{}
    data1.Insert(2)
    data1.Insert(1)
    if data1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
    
    val,err := data1.Get(0)
    if err != nil {
        t.Error("illegal data get error.")
    }
    if val != 2 {
        t.Error("illegal data get.")
    }
    
    //範囲外の指定
    val,err = data1.Get(10)
    if err == nil {
        t.Error("illegal error.")
    }
    if err.Error() != "out of range." {
        t.Error("illegal error message.")
    }
}
func TestArrayIntegerSearch(t *testing.T) {
    var data1  = ArrayInteger{}
    data1.Insert(1)
    data1.Insert(2)
    data1.Insert(1)
    res := data1.Search(1)
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
func TestArrayIntegerSearchNoMatch(t *testing.T) {
    var data1  = ArrayInteger{}
    data1.Insert(1)
    data1.Insert(2)
    data1.Insert(1)
    res := data1.Search(3)
    if len(res) != 0 {
        t.Error("illegal result len.")
    }
}
