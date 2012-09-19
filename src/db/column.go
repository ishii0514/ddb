/*
 dbパッケージ
*/
package db

import (

)


//カラムインターフェース
type Column interface {
	Name() string
}

//INT型のカラム
//TODO intをInteger型に置き換え
type ColumnInteger struct{
    name string
    data []Integer
}

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
	insertValue := INVALID_VALUE_INTEGER
	v,err := StringtoInteger(data)
	if err == nil {
		insertValue = v			
    }
	p.Insert(insertValue)
}

