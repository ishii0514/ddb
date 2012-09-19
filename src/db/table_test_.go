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
func TestgetInsertVaue(t *testing.T) {
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

//insertのテスト 3カラム
/*
func TestTableInsert3columns(t *testing.T) {
    var table1 = createTable(3)
    table1.Insert([]string{"1", "10","100"})
    table1.Insert([]string{"2", "20","200"})
    table1.Insert([]string{"3", "30","300"})
    
    if table1.DataCount() != 3 {
        t.Error("illegal data count.")
    }
}
*/