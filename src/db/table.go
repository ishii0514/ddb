package db

import (
)
//TODO カラムの生成
//TODO テーブルの生成
//TODO Insertのテスト
//テーブル
type Table struct{
    name string
    datacount ROWNUM
    columns []*ColumnInteger
}

//テーブル名の取得
func (p *Table) Name() string { return p.name}

//データ数
func (p *Table) DataCount() ROWNUM { return p.datacount}
/**
* データ挿入
* カラム数より少ない場合、変換できない値が入ってる場合、0で埋められる。
* カラム数より多い列のデータは無視される。
*/
func (p *Table) Insert(values []string) {
	//カラム毎にinsert
	if len(p.columns) == 0{
		return
	}
	for i, column := range p.columns {
	    column.InsertByString(getInsertVaue(i,&values))
	}
	p.datacount += 1
}

//投入データを決定する。
func getInsertVaue(columnIndex int,insertValues *[]string) string{
	insertValue := INVALID_VALUE
	//カラムインデックスの方が大きい場合
	if columnIndex >= len(*insertValues){
		return insertValue
	}
	return (*insertValues)[columnIndex]
}
//テーブルを生成する
/*
func createTable(tablename string,columnNumber int) *Table{
	newTable := Table{name : tablename}
	
	return &newTable
}
*/
