package db

import (
)
//TODO 入ってるデータの確認
//TODO データフェッチ
//TODO 一括ロード
//TODO Search
//TODO Insert
//TODO Delete
//TODO テーブルの生成 ファクトリメソッド

//テーブル
type Table struct{
    name string
    datacount ROWNUM
    columns []Column
}

//テーブル名の取得
func (p *Table) Name() string { return p.name}

//データ数
func (p *Table) DataCount() ROWNUM { return p.datacount}


/**
* データ挿入
* カラム数より少ない場合、変換できない値が入ってる場合、INVALID_VALUEで埋められる。
* カラム数より多い列のデータは無視される。
*/
func (p *Table) Insert(values []string) {
	if len(p.columns) == 0{
		return
	}
	//カラム毎にinsert
	for i, column := range p.columns {
	    column.Insert(getInsertVaue(i,&values))
	}
	p.datacount += 1
}

//投入データを決定する。
func getInsertVaue(columnIndex int,insertValues *[]string) string{
	//カラムインデックスの方が大きい場合
	if columnIndex >= len(*insertValues){
		return INVALID_VALUE
	}
	return (*insertValues)[columnIndex]
}
/**
* カラムの追加
* データ件数が0件の場合のみ
*/
func (p *Table) AddColumn(name string,columntype ColumnType) error {
	if p.DataCount() >0 {
		//データ既にある場合
		return &DbError{"data exisits."}
	}
	column := createColumn(name,columntype)
	if column == nil {
		return &DbError{"add column failed."}
	}
	p.columns = append(p.columns,column)
	return nil
}
/**
* カラムの数
*/
func (p *Table) ColumnCount() int { 
	return len(p.columns)
}

//テーブルを生成する
/*
func createTable(tablename string,columnNumber int) *Table{
	newTable := Table{name : tablename}
	
	return &newTable
}
*/
