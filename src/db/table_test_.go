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

//insertのテスト
func TestTableInsert(t *testing.T) {
    var table1 = Table{name : "tab1"}
    table1.Insert([]string{"1", "10","100"})
    table1.Insert([]string{"2", "20","200"})
    table1.Insert([]string{"3", "30","300"})
    
    if table1.DataCount() != 3 {
        t.Error("illegal table name.")
    }
}