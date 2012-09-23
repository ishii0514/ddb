package db

import (

)

//INTEGER型のデータ構造インターフェース
type DataInteger interface {
	DataCount() ROWNUM
	Get(ROWNUM) (Integer,error)
	Search(Integer) []ROWNUM
	Insert(Integer)
}

//Arry型のデータ構造
type ArrayInteger struct{
    data []Integer
}

//データ数の取得
func (p *ArrayInteger) DataCount() ROWNUM { return ROWNUM(len(p.data))}

//指定した行のデータを取得
func (p *ArrayInteger) Get(row ROWNUM) (Integer,error) {
    if row >= p.DataCount(){
        return 0,&DbError{"out of range."}
    }
	return p.data[row],nil
}
//指定した値の行リストを返す
func (p *ArrayInteger) Search(searchValue Integer) []ROWNUM {
	res := make([]ROWNUM,0)
	for i, v := range p.data {
		if searchValue == v {
			res = append(res,ROWNUM(i))
		}
	}
	return res
}
//データ挿入
func (p *ArrayInteger) Insert(data Integer) {
	p.data = append(p.data,data)
}
