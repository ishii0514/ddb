package db

import (

)
//TODO 範囲検索

//INTEGER型のデータ構造インターフェース
type DataInteger interface {
  DataCount() ROWNUM
  Get(ROWNUM) (Integer,error)
  Search(Integer) []ROWNUM
  Insert(Integer) ROWNUM
  Delete(Integer) ROWNUM
}

//Array型のデータ構造
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
/**
 * データ挿入
 * インサート件数を返す
 */
func (p *ArrayInteger) Insert(data Integer) ROWNUM{
  p.data = append(p.data,data)
  return ROWNUM(1)
}
/**
 * データ削除
 * 削除件数を返す
 */
func (p *ArrayInteger) Delete(deleteValue Integer) ROWNUM{
  rows := p.Search(deleteValue)
  for i, row := range rows {
    //削除分を考慮する
    p.delete(row-ROWNUM(i))
  }
  return ROWNUM(len(rows))
}
//行を指定してデータ削除
func (p *ArrayInteger) delete(deleteROW ROWNUM){
  p.data = append(p.data[:deleteROW],p.data[deleteROW+1:]...)
  c := make([]Integer,len(p.data))
  copy(c,p.data)
  p.data = c
}
