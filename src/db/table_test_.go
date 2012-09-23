package db

import (
    "testing"
)

func TestTableName(t *testing.T) {
    var table1 = Table{name : "tab1"}
    if table1.Name() != "tab1" {
        t.Error("illegal table name.")
    }
}

//insertのテスト 0カラム
func TestTableInsert0column(t *testing.T) {
    var table1 = Table{name : "tab1"}
    table1.Insert([]string{"1", "10","100"})
    table1.Insert([]string{"2", "20","200"})
    table1.Insert([]string{"3", "30","300"})
    
    if table1.DataCount() != 0 {
        t.Error("illegal data count.")
    }
}

//追加データの判定
func TestGetInsertVaue(t *testing.T) {
	values:= []string{"1", "10","STR"}
	
	//２番目の項目
	res := getInsertVaue(1,&values)
	if res != "10" {
        t.Error("illegal insert data.")
    }
    
    //カラム数を超えた場合
    res = getInsertVaue(3,&values)
	if res != INVALID_VALUE {
        t.Error("illegal insert data.")
    }
}
//カラムの追加
func TestTableAddcolumns(t *testing.T) {
    var table1 = Table{name : "tab1"}
    if table1.ColumnCount() != 0 {
        t.Error("illegal column count.")
    }
    
    err := table1.AddColumn("col1",COLUMN_TYPE_INTEGER)
    if err != nil {
        t.Error("add column failed.#1")
    }
    err = table1.AddColumn("col2",COLUMN_TYPE_STRING)
    if err == nil {
        t.Error("illegal column added.")
    }
    err = table1.AddColumn("col3",COLUMN_TYPE_INTEGER)
    if err != nil {
        t.Error("add column failed.#2")
    }
    
    if table1.ColumnCount() != 2 {
        t.Error("illegal column count.")
    }
}

//insertのテスト 3カラム
func TestTableInsert3columns(t *testing.T) {
    table1 := Table{name : "tab1"}
    table1.AddColumn("col1",COLUMN_TYPE_INTEGER)
    table1.AddColumn("col2",COLUMN_TYPE_INTEGER)
    table1.AddColumn("col3",COLUMN_TYPE_INTEGER)
    
    
    table1.Insert([]string{"1", "10","100"})
    table1.Insert([]string{"2", "20","200"})
    table1.Insert([]string{"3", "30","300"})
    table1.Insert([]string{"4", "40","400","4000"})
    table1.Insert([]string{"5", "str","500"})
    
    if table1.DataCount() != 5 {
        t.Error("illegal data count.")
    }
}