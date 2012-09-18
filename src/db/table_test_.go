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