/*
 dbパッケージ
*/
package db

import (

)


//カラムインターフェース
type Column interface {
	Type() ColumnType
	Name() string
	InsertByString(data string)
}
//カラムの作成
func createColumn(columnName string, columnType ColumnType) Column{
	switch {
    case columnType == COLUMN_TYPE_INTEGER:
    	newColumn := ColumnInteger{name : columnName}
    	return &newColumn
    }
    return nil
}



//INT型のカラム
type ColumnInteger struct{
    name string
    data []Integer
}

//カラムタイプ
func (p *ColumnInteger) Type() ColumnType { return COLUMN_TYPE_INTEGER}
//カラム名の取得
func (p *ColumnInteger) Name() string { return p.name}

//データ数の取得
func (p *ColumnInteger) DataCount() ROWNUM { return ROWNUM(len(p.data))}

//指定した行のデータを取得
func (p *ColumnInteger) Get(row ROWNUM) (Integer,error) {
    if row >= p.DataCount(){
        return 0,&DbError{"out of range."}
    }
	return p.data[row],nil
}

//指定した値の行リストを返す
func (p *ColumnInteger) Search(searchValue Integer) []ROWNUM {
	res := make([]ROWNUM,0)
	for i, v := range p.data {
		res = appendData(searchValue==v,ROWNUM(i),res)
	}
	return res
}

//データ挿入
func (p *ColumnInteger) Insert(data Integer) {
	p.data = append(p.data,data)
}

//指定した行のデータを削除する
func (p *ColumnInteger) DeleteAt(row ROWNUM) {
    if row >= p.DataCount(){
        return
    }
    p.data = append(p.data[:row],p.data[row+1:]...)
}

//条件に一致したらデータを追加する
//TODO 使われ方が変なのであとでリファクタ
func appendData(isAppend bool ,value ROWNUM, values []ROWNUM) []ROWNUM{
	if isAppend {
		values = append(values,value)
	}
	return values
}

/**
 * データ挿入　文字列入力
 * 不正なデータの場合、INVALID_VALUE_INTEGERを挿入する。
 */
func (p *ColumnInteger) InsertByString(data string) {
	//型チェック
	p.Insert(convertToInteger(data))
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