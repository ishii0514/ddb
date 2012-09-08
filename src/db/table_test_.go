package db

import (
    "testing"
)

func TestTableName(t *testing.T) {
    var table1 = Table{name : "tab1"}
    if table1.GetName() != "tab1" {
        t.Error("illegal table name.")
    }
}
