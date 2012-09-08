package db

import (
	
)

//テーブル
type Table struct{
    name string
    columns []*ColumnInteger
}

//テーブル名の取得
func (p *Table) GetName() string { return p.name}