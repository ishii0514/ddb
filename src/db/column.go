/*
 dbパッケージ
*/
package db

import (
)

//エラー構造体
type DbError struct {
    Message string
}
func (e *DbError) Error() string {
    return e.Message
}


//カラムインターフェース
type Column interface {
	GetName() string
}

//INT型のカラム
type ColumnInteger struct{
    name string
    data []int
}

//カラム名の取得
func (p *ColumnInteger) GetName() string { return p.name}

//データ数の取得
func (p *ColumnInteger) DataCount() int { return len(p.data)}

//指定した行のデータを取得
func (p *ColumnInteger) Get(row int) (int,error) {
    if row >= p.DataCount(){
        return 0,&DbError{"out of range."}
    }
	return p.data[row],nil
}
//指定した値の行を返す
func (p *ColumnInteger) Search(searchValue int) []int {
	res := make([]int,0)
	for i, v := range p.data {
		if v== searchValue {
			res = append(res,i)
		}
	}
	return res
}

//データ挿入
func (p *ColumnInteger) Insert(data int) {p.data = append(p.data,data)}

//指定した行のデータを削除する
func (p *ColumnInteger) DeleteAt(row int) {
    if row >= p.DataCount(){
        return
    }
    p.data = append(p.data[:row],p.data[row+1:]...)
}