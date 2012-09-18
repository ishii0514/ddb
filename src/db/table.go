package db

import (
	"strconv"
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

//指定された数の項目を作成する
func (p *Table) createColumns(number int) {
	
}

/**
* データ挿入
* カラム数より少ない場合、変換できない値が入ってる場合、0で埋められる。
* カラム数より多い列のデータは無視される。
*/
func (p *Table) Insert(values []string) {
	//カラム毎にinsert
	//TODO 並列化できるかも
	if len(p.columns) == 0{
		return
	}
	for i, column := range p.columns {
	    column.Insert(getInsertVaue(i,&values))
	}
	p.datacount += 1
}

//投入データを決定する。
func getInsertVaue(columnIndex int,insertValues *[]string) int{
	INVALID_VALUE := 0
	insertValue := INVALID_VALUE
	
	//カラムインデックスの方が大きい場合
	if columnIndex >= len(*insertValues){
		return insertValue
	}
	//型変換
	v,err := strconv.Atoi((*insertValues)[columnIndex])
	if err == nil {
		insertValue = v			
    }
	return insertValue
}
//テーブルを生成する
/*
func createTable(tablename string,columnNumber int) *Table{
	newTable := Table{name : tablename}
	
	return &newTable
}
*/