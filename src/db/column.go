/*
 dbパッケージ
*/
package db

import (

)
//TODO Search

//カラムインターフェース
type Column interface {
	Type() ColumnType
	Name() string
	DataCount() ROWNUM
	Search(string) []ROWNUM
	Insert(string) ROWNUM
	Delete(string) ROWNUM
}

//カラムの生成
func createColumn(columnName string, columnType ColumnType) Column{
	switch {
    case columnType == COLUMN_TYPE_INTEGER:
    	newColumn := ColumnInteger{name : columnName,data:new(ArrayInteger)}
    	return &newColumn
    }
    return nil
}


//INT型のカラム
type ColumnInteger struct{
    name string
    data DataInteger
}
//カラムタイプ
func (p *ColumnInteger) Type() ColumnType { return COLUMN_TYPE_INTEGER}
//カラム名の取得
func (p *ColumnInteger) Name() string { return p.name}

//データ数の取得
func (p *ColumnInteger) DataCount() ROWNUM { return p.data.DataCount()}

//指定した行のデータを取得
func (p *ColumnInteger) Get(row ROWNUM) (Integer,error) {
    return p.data.Get(row)
}

//指定した値の行リストを返す
func (p *ColumnInteger) Search(searchValue string) []ROWNUM {
	//型チェック
	v,err := StringtoInteger(searchValue)
	if err == nil {
		return p.data.Search(v)
    }
	return []ROWNUM{}
}

/**
 * データ挿入　文字列入力
 * 不正なデータの場合、INVALID_VALUE_INTEGERを挿入する。
 */
func (p *ColumnInteger) Insert(data string)  ROWNUM{
	//型チェックしてインサート
	return p.data.Insert(convertToInteger(data))
}


// 文字列入力に対して型チェックとコンバートを行う
// Integer型に変換できない場合、INVALID_VALUE_INTEGERを返す
func convertToInteger(data string) Integer{
	//無効値
	if data == INVALID_VALUE {
		return INVALID_VALUE_INTEGER
	}
	//型変換
	v,err := StringtoInteger(data)
	if err == nil {
		return v
    }
    //変換に失敗したら無効値
    return INVALID_VALUE_INTEGER
}
/**
 * データ削除
 */
func (p *ColumnInteger) Delete(data string)  ROWNUM{
	//型チェック
	v,err := StringtoInteger(data)
	if err == nil {
		return p.data.Delete(v)
    }
    return 0
}