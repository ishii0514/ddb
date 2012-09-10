package db

import (
	
)

//テーブル
type Table struct{
    name string
    datacount int
    columns []*ColumnInteger
}

//テーブル名の取得
func (p *Table) Name() string { return p.name}

//データ数
func (p *Table) DataCount() int { return p.datacount}

//データ挿入
func (p *Table) Insert(values []string) {

	p.datacount += 1
}