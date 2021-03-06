package db

import (
    "testing"
)


func TestColumnInteger(t *testing.T) {
    var col1 = ColumnInteger{name : "col1"}
    if col1.Name() != "col1" {
        t.Error("illegal name.")
    }
}

func TestColumnIntegerInsert(t *testing.T) {
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("1")
    if col1.DataCount() != 1 {
        t.Error("illegal datacount.")
    }
    col1.Insert("1")
    if col1.DataCount() != 2 {
        t.Error("illegal datacount.")
    }
}
//文字列INSERTのテスト
func TestColumnInsertIllegalData(t *testing.T) {
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("1")
    col1.Insert("nointeger")
    col1.Insert("2")

    if len(col1.Search("1")) != 1 {
        t.Error("illegal result len.")
    }
    if col1.DataCount() != 3{
    	t.Error("illegal datacount.")
    }
}
func TestColumnIntegerGet(t *testing.T) {
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("2")
    col1.Insert("1")
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
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("1")
    col1.Insert("2")
    col1.Insert("1")
    res := col1.Search("1")
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
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("1")
    col1.Insert("2")
    col1.Insert("1")
    res := col1.Search("3")
    if len(res) != 0 {
        t.Error("illegal result len.")
    }
}
func TestColumnIntegerSearchIllegalNum(t *testing.T) {
    var col1  = ColumnInteger{name : "col1",data:new(ArrayInteger)}
    col1.Insert("1")
    col1.Insert("2")
    col1.Insert("1")
    res := col1.Search("str")
    if len(res) != 0 {
        t.Error("illegal result len.")
    }
}
//削除
func TestColumnIntegerDelete(t *testing.T) {
	col1 := createColumn("columnA",COLUMN_TYPE_INTEGER)
	col1.Insert("1")
    col1.Insert("2")
    col1.Insert("1")
    col1.Insert("1")
    if col1.DataCount() != 4 {
        t.Error("illegal data count.")
    }
    
    //削除
    delCnt := col1.Delete("1")
    if col1.DataCount() != 1 {
        t.Error("illegal delete.")
    }
    if delCnt != 3 {
        t.Error("illegal delete count.")
    }
    
    //削除 該当する値なし
    delCnt = col1.Delete("1")
    if col1.DataCount() != 1 {
        t.Error("illegal no delete.")
    }
    if delCnt != 0 {
        t.Error("illegal no delete count.")
    }
}
//削除
func TestColumnIntegerDeleteIllegalNum(t *testing.T) {
	col1 := createColumn("columnA",COLUMN_TYPE_INTEGER)
	col1.Insert("1")
    col1.Insert("2")
    col1.Insert("1")
    col1.Insert("1")
    if col1.DataCount() != 4 {
        t.Error("illegal data count.")
    }
    
    //文字列による削除
    delCnt := col1.Delete("str")
    if col1.DataCount() != 4 {
        t.Error("illegal delete.")
    }
    if delCnt != 0 {
        t.Error("illegal delete count.")
    }
    
    //文字列による削除
    delCnt = col1.Delete("001")
    if col1.DataCount() != 1 {
        t.Error("illegal delete.")
    }
    if delCnt != 3 {
        t.Error("illegal delete count.")
    }

}
//データ変換のテスト
func TestColumnConvertToInteger(t *testing.T) {
    if convertToInteger(INVALID_VALUE) != INVALID_VALUE_INTEGER {
        t.Error("illegal convert.#1")
    }
    if convertToInteger("999") != 999 {
        t.Error("illegal convert.#2")
    }
    if convertToInteger("helloworld") != INVALID_VALUE_INTEGER {
        t.Error("illegal convert.#3")
    }
}
//カラムの作成
func TestColumnCreateColumn(t *testing.T) {
	column := createColumn("columnA",COLUMN_TYPE_INTEGER)
	if column == nil{
		t.Error("illegal column create.#1")
	}
	if column.Name() != "columnA"{
		t.Error("illegal column create.#2")
	}
	
	column = createColumn("columnA",COLUMN_TYPE_STRING)
	if column != nil{
		t.Error("illegal column create.#3")
	}
}

